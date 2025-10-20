import { Message } from 'tdesign-vue-next';

let errorLock = false;
function showError(msg) {
  if (!errorLock) {
    errorLock = true;
    Message.error(msg);
    setTimeout(() => {
      errorLock = false;
    }, 1000);
  }
}

export default function registerInterceptors(http) {
  // 请求拦截器
  http.interceptors.request.use(
    (originConf) => {
      const config = { ...originConf };
      const token = localStorage.getItem('token');
      if (token) {
        config.headers.Authorization = `Bearer ${token}`;
      }
      // 业务错误弹窗控制，默认true
      if (typeof config.showBizError === 'undefined') {
        config.showBizError = true;
      }
      // 统一加上共同参数
      const commonParams = { appId: 'yourAppId', platform: 'web' };
      if (config.method === 'get') {
        config.params = { ...commonParams, ...(config.params || {}) };
      } else if (['post', 'put', 'patch'].includes(config.method)) {
        if (config.data instanceof FormData) {
          Object.entries(commonParams).forEach(([k, v]) => config.data.append(k, v));
        } else {
          config.data = { ...commonParams, ...(config.data || {}) };
        }
      }
      return config;
    },
    (error) => {
      showError('请求配置异常');
      return Promise.reject({
        isBizError: false,
        code: undefined,
        message: '请求配置异常',
        raw: error,
      });
    },
  );

  // 响应拦截器
  http.interceptors.response.use(
    (response) => {
      const { data } = response;
      const showErrorFlag = response.config && response.config.showBizError !== false;
      if (typeof data === 'object' && data !== null && 'code' in data) {
        if (data.code === 0) {
          return data;
        }
        if (data.code === 401) {
          localStorage.removeItem('token');
          showError('未登录，请重新登录');
          window.location.href = '/login';
        } else if (showErrorFlag) {
          showError(data.message || '业务错误');
        }
        return Promise.reject({
          isBizError: true,
          code: data.code,
          message: data.message || '业务错误',
          raw: data,
        });
      }
      if (showErrorFlag) {
        showError('接口返回格式错误');
      }
      return Promise.reject({
        isBizError: true,
        code: undefined,
        message: '接口返回格式错误',
        raw: data,
      });
    },
    (error) => {
      const { response } = error;
      if (response) {
        const code = response.status;
        let msg = '';
        switch (code) {
          case 401:
            localStorage.removeItem('token');
            window.location.href = '/login';
            msg = '未授权，请重新登录';
            break;
          case 403:
            msg = '无权限访问';
            break;
          case 500:
            msg = '服务器错误，请稍后重试';
            break;
          default:
            msg = response.data?.message || `请求出错（${code}）`;
        }
        showError(msg);
        return Promise.reject({
          isBizError: false,
          code,
          message: msg,
          raw: error,
        });
      }
      showError('网络连接异常');
      return Promise.reject({
        isBizError: false,
        code: undefined,
        message: '网络连接异常',
        raw: error,
      });
    },
  );
}
