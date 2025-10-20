// 大文件上传类
class LargeFileUploader {
    constructor() {
        this.chunkSize = 2 * 1024 * 1024; // 2MB 分片
        this.maxConcurrent = 3; // 最大并发数
        this.uploadQueue = [];
        this.uploading = false;
        this.currentUploads = 0; // 当前正在上传的数量
        this.paused = false;
        this.uploadedChunks = new Set();
        this.failedChunks = new Set();
        this.onProgress = null; // 进度回调
        this.onComplete = null; // 完成回调
        this.onError = null; // 错误回调
    }

    // 设置回调函数
    setCallbacks(onProgress, onComplete, onError) {
        this.onProgress = onProgress;
        this.onComplete = onComplete;
        this.onError = onError;
    }

    // 文件分片上传 - 修正后的并发控制
    async uploadLargeFile(file) {
        const fileId = this.generateFileId();
        const chunks = this.splitFileIntoChunks(file);
        
        console.log(`开始上传文件: ${file.name}, 大小: ${this.formatFileSize(file.size)}, 分片数: ${chunks.length}`);
        
        // 1. 初始化上传
        const initResult = await this.initUpload(fileId, file.name, file.size);
        if (initResult.code !== 200) {
            throw new Error(initResult.message || '初始化上传失败');
        }

        // 2. 创建上传任务队列
        this.uploadQueue = chunks.map((chunk, index) => ({
            fileId,
            chunkIndex: index,
            chunk,
            retryCount: 0
        }));

        // 3. 开始并发上传
        await this.startConcurrentUpload();

        // 4. 完成上传
        return await this.completeUpload(fileId);
    }

    // 开始并发上传 - 真正的并发控制
    async startConcurrentUpload() {
        this.uploading = true;
        this.paused = false;
        this.uploadedChunks.clear();
        this.failedChunks.clear();
        
        const uploadPromises = [];
        
        // 创建上传工作函数 - 只负责任务调度
        const uploadWorker = async () => {
            while (this.hasMoreTasks() && !this.paused) {
                // 等待并发槽位可用
                await this.waitForConcurrencySlot();
                
                const task = this.getNextTask();
                if (!task) break;

                // 执行单个任务
                await this.executeTask(task);
            }
        };

        // 启动多个工作线程
        for (let i = 0; i < this.maxConcurrent; i++) {
            uploadPromises.push(uploadWorker());
        }

        // 等待所有上传完成
        await Promise.all(uploadPromises);
        
        // 检查是否有失败的分片
        if (this.failedChunks.size > 0) {
            throw new Error(`上传失败的分片: ${Array.from(this.failedChunks).join(', ')}`);
        }
    }

    // 检查是否还有更多任务
    hasMoreTasks() {
        return this.uploadQueue.length > 0;
    }

    // 等待并发槽位可用
    async waitForConcurrencySlot() {
        while (this.currentUploads >= this.maxConcurrent) {
            await this.sleep(100);
        }
    }

    // 获取下一个任务
    getNextTask() {
        return this.uploadQueue.shift();
    }

    // 执行单个任务 - 只负责任务执行和结果处理
    async executeTask(task) {
        this.currentUploads++;
        
        try {
            await this.uploadChunkWithRetry(task);
            this.onTaskSuccess(task);
        } catch (error) {
            this.onTaskError(task, error);
        } finally {
            this.currentUploads--;
        }
    }

    // 任务成功处理
    onTaskSuccess(task) {
        this.uploadedChunks.add(task.chunkIndex);
        this.updateProgress();
    }

    // 任务错误处理
    onTaskError(task, error) {
        console.error(`分片 ${task.chunkIndex} 上传失败:`, error);
        this.failedChunks.add(task.chunkIndex);
        
        // 处理重试逻辑
        this.handleRetry(task);
    }

    // 处理重试逻辑
    handleRetry(task) {
        if (this.shouldRetry(task)) {
            task.retryCount++;
            this.uploadQueue.unshift(task); // 重新加入队列头部
        }
    }

    // 判断是否应该重试
    shouldRetry(task) {
        return task.retryCount < 3;
    }

    // 带重试的分片上传
    async uploadChunkWithRetry(task) {
        const maxRetries = 3;
        let lastError;

        for (let attempt = 0; attempt <= maxRetries; attempt++) {
            try {
                return await this.uploadChunk(task);
            } catch (error) {
                lastError = error;
                console.warn(`分片 ${task.chunkIndex} 第 ${attempt + 1} 次尝试失败:`, error);
                
                if (attempt < maxRetries) {
                    // 指数退避重试
                    await this.waitBeforeRetry(attempt);
                }
            }
        }

        throw lastError;
    }

    // 上传单个分片 - 纯上传逻辑
    async uploadChunk(task) {
        const formData = this.createChunkFormData(task);
        const response = await this.sendChunkRequest(formData);
        return this.processChunkResponse(response);
    }

    // 创建分片表单数据
    createChunkFormData(task) {
        const formData = new FormData();
        formData.append('file', task.chunk);
        formData.append('fileId', task.fileId);
        formData.append('chunkIndex', task.chunkIndex);
        return formData;
    }

    // 发送分片请求
    async sendChunkRequest(formData) {
        const response = await fetch(`/api/upload/chunk`, {
            method: 'POST',
            body: formData
        });

        if (!response.ok) {
            throw new Error(`HTTP ${response.status}: ${response.statusText}`);
        }

        return response;
    }

    // 处理分片响应
    async processChunkResponse(response) {
        const result = await response.json();
        if (result.code !== 200) {
            throw new Error(result.message || '上传失败');
        }
        return result.data;
    }

    // 重试前等待
    async waitBeforeRetry(attempt) {
        const delay = Math.pow(2, attempt) * 1000;
        await this.sleep(delay);
    }

    // 暂停上传
    pauseUpload() {
        this.paused = true;
        console.log('上传已暂停');
    }

    // 恢复上传
    resumeUpload() {
        this.paused = false;
        console.log('上传已恢复');
    }

    // 取消上传
    cancelUpload() {
        this.paused = true;
        this.uploadQueue = [];
        this.uploading = false;
        console.log('上传已取消');
    }

    // 更新进度
    updateProgress() {
        const totalChunks = this.uploadedChunks.size + this.failedChunks.size + this.uploadQueue.length;
        const progress = totalChunks > 0 ? (this.uploadedChunks.size / totalChunks) * 100 : 0;
        
        // 调用进度回调
        if (this.onProgress) {
            this.onProgress({
                progress: Math.round(progress),
                uploadedChunks: this.uploadedChunks.size,
                failedChunks: this.failedChunks.size,
                queuedChunks: this.uploadQueue.length,
                currentConcurrent: this.currentUploads,
                totalChunks: totalChunks
            });
        }

        console.log(`上传进度: ${Math.round(progress)}% (${this.uploadedChunks.size}/${totalChunks})`);
    }

    // 工具方法
    sleep(ms) {
        return new Promise(resolve => setTimeout(resolve, ms));
    }

    formatFileSize(bytes) {
        if (bytes === 0) return '0 Bytes';
        const k = 1024;
        const sizes = ['Bytes', 'KB', 'MB', 'GB'];
        const i = Math.floor(Math.log(bytes) / Math.log(k));
        return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
    }

    // 文件分片
    splitFileIntoChunks(file) {
        const chunks = [];
        let start = 0;
        
        while (start < file.size) {
            const end = Math.min(start + this.chunkSize, file.size);
            chunks.push(file.slice(start, end));
            start = end;
        }
        
        return chunks;
    }

    // 初始化上传
    async initUpload(fileId, fileName, fileSize) {
        const response = await fetch('/api/upload/init', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                fileId,
                fileName,
                fileSize,   
                chunkSize: this.chunkSize
            })
        });

        const result = await response.json();
        if (result.code !== 200) {
            throw new Error(result.message || '初始化上传失败');
        }
        return result;
    }

    // 完成上传
    async completeUpload(fileId) {
        const response = await fetch('/api/upload/complete', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ fileId })
        });

        const result = await response.json();
        if (result.code !== 200) {
            throw new Error(result.message || '完成上传失败');
        }
        return result.data;
    }

    // 生成文件ID
    generateFileId() {
        return Date.now() + '_' + Math.random().toString(36).substr(2, 9);
    }

    // 获取上传状态
    getUploadStatus() {
        return {
            uploading: this.uploading,
            paused: this.paused,
            currentUploads: this.currentUploads,
            uploadedChunks: this.uploadedChunks.size,
            failedChunks: this.failedChunks.size,
            queuedChunks: this.uploadQueue.length
        };
    }
}

// 导出供其他模块使用
window.LargeFileUploader = LargeFileUploader; 