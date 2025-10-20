import axios from 'axios';
import registerInterceptors from './interceptors';

const http = axios.create({
  baseURL: '/', // 可根据需要修改
  timeout: 10000,
});

registerInterceptors(http);

export default http;
