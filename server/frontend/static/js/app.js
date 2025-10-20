// 前端应用主文件
class App {
    constructor() {
        this.apiBase = '/api';
        this.init();
    }

    init() {
        this.bindEvents();
        this.loadUserInfo();
    }

    bindEvents() {
        // 绑定表单提交事件
        document.addEventListener('DOMContentLoaded', () => {
            this.bindFormEvents();
            this.bindNavigationEvents();
        });
    }

    bindFormEvents() {
        // 用户登录表单
        const loginForm = document.getElementById('loginForm');
        if (loginForm) {
            loginForm.addEventListener('submit', (e) => {
                e.preventDefault();
                this.handleLogin();
            });
        }

        // 用户注册表单
        const registerForm = document.getElementById('registerForm');
        if (registerForm) {
            registerForm.addEventListener('submit', (e) => {
                e.preventDefault();
                this.handleRegister();
            });
        }

        // 微信urlLink生成表单
        const wechatForm = document.getElementById('wechatForm');
        if (wechatForm) {
            wechatForm.addEventListener('submit', (e) => {
                e.preventDefault();
                this.handleWechatUrlLink();
            });
        }
    }

    bindNavigationEvents() {
        // 导航链接点击事件
        const navLinks = document.querySelectorAll('.navbar-nav a');
        navLinks.forEach(link => {
            link.addEventListener('click', (e) => {
                const href = link.getAttribute('href');
                if (href && href.startsWith('#')) {
                    e.preventDefault();
                    this.navigateTo(href.substring(1));
                }
            });
        });
    }

    async loadUserInfo() {
        try {
            const response = await this.apiCall('GET', '/user/profile');
            if (response.success) {
                this.updateUserInfo(response.data);
            }
        } catch (error) {
            console.log('用户未登录或获取用户信息失败');
        }
    }

    updateUserInfo(user) {
        const userInfoElement = document.getElementById('userInfo');
        if (userInfoElement && user) {
            userInfoElement.innerHTML = `
                <div class="alert alert-info">
                    <strong>当前用户:</strong> ${user.username} (${user.email})
                </div>
            `;
        }
    }

    async handleLogin() {
        const form = document.getElementById('loginForm');
        const formData = new FormData(form);
        
        const loginData = {
            username: formData.get('username'),
            password: formData.get('password')
        };

        try {
            const response = await this.apiCall('POST', '/user/login', loginData);
            if (response.success) {
                this.showAlert('登录成功！', 'success');
                this.loadUserInfo();
                setTimeout(() => {
                    window.location.href = '/';
                }, 1000);
            } else {
                this.showAlert(response.message || '登录失败', 'danger');
            }
        } catch (error) {
            this.showAlert('登录失败: ' + error.message, 'danger');
        }
    }

    async handleRegister() {
        const form = document.getElementById('registerForm');
        const formData = new FormData(form);
        
        const registerData = {
            username: formData.get('username'),
            email: formData.get('email'),
            password: formData.get('password')
        };

        try {
            const response = await this.apiCall('POST', '/user/register', registerData);
            if (response.success) {
                this.showAlert('注册成功！', 'success');
                setTimeout(() => {
                    window.location.href = '/login';
                }, 1000);
            } else {
                this.showAlert(response.message || '注册失败', 'danger');
            }
        } catch (error) {
            this.showAlert('注册失败: ' + error.message, 'danger');
        }
    }

    async handleWechatUrlLink() {
        const form = document.getElementById('wechatForm');
        const formData = new FormData(form);
        
        const urlLinkData = {
            path: formData.get('path'),
            query: formData.get('query'),
            is_expire: formData.get('is_expire') === 'on',
            expire_type: parseInt(formData.get('expire_type')),
            expire_time: formData.get('expire_time') ? parseInt(formData.get('expire_time')) : 0,
            expire_interval: formData.get('expire_interval') ? parseInt(formData.get('expire_interval')) : 0
        };

        try {
            const response = await this.apiCall('POST', '/wechat/url-link', urlLinkData);
            if (response.success) {
                this.showAlert('urlLink生成成功！', 'success');
                this.displayWechatResult(response.data);
            } else {
                this.showAlert(response.message || '生成失败', 'danger');
            }
        } catch (error) {
            this.showAlert('生成失败: ' + error.message, 'danger');
        }
    }

    displayWechatResult(data) {
        const resultElement = document.getElementById('wechatResult');
        if (resultElement && data) {
            resultElement.innerHTML = `
                <div class="alert alert-success">
                    <h4>生成结果</h4>
                    <p><strong>urlLink:</strong> <a href="${data.url_link}" target="_blank">${data.url_link}</a></p>
                    <p><strong>状态:</strong> ${data.status}</p>
                    <p><strong>创建时间:</strong> ${new Date(data.create_time * 1000).toLocaleString()}</p>
                </div>
            `;
        }
    }

    async apiCall(method, endpoint, data = null) {
        const url = this.apiBase + endpoint;
        const options = {
            method: method,
            headers: {
                'Content-Type': 'application/json',
            }
        };

        if (data) {
            options.body = JSON.stringify(data);
        }

        try {
            const response = await fetch(url, options);
            const result = await response.json();
            
            if (!response.ok) {
                throw new Error(result.message || '请求失败');
            }
            
            return result;
        } catch (error) {
            console.error('API调用失败:', error);
            throw error;
        }
    }

    showAlert(message, type = 'info') {
        const alertContainer = document.getElementById('alertContainer');
        if (alertContainer) {
            const alertElement = document.createElement('div');
            alertElement.className = `alert alert-${type}`;
            alertElement.textContent = message;
            
            alertContainer.appendChild(alertElement);
            
            // 3秒后自动移除
            setTimeout(() => {
                alertElement.remove();
            }, 3000);
        }
    }

    navigateTo(page) {
        // 简单的页面导航逻辑
        const pages = document.querySelectorAll('.page');
        pages.forEach(p => p.style.display = 'none');
        
        const targetPage = document.getElementById(page);
        if (targetPage) {
            targetPage.style.display = 'block';
        }
    }

    // 工具方法
    formatDate(timestamp) {
        return new Date(timestamp * 1000).toLocaleString();
    }

    copyToClipboard(text) {
        navigator.clipboard.writeText(text).then(() => {
            this.showAlert('已复制到剪贴板', 'success');
        }).catch(() => {
            this.showAlert('复制失败', 'danger');
        });
    }
}

// 初始化应用
const app = new App();

// 导出供其他模块使用
window.App = App;
window.app = app; 