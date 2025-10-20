// 用户信息 Store
import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', {
  state: () => ({
    userInfo: {
      name: '用户',
      avatar: '',
      email: '',
      phone: ''
    },
    isLoggedIn: false
  }),
  
  getters: {
    getUserName: (state) => state.userInfo.name,
    getUserAvatar: (state) => state.userInfo.avatar,
    isUserLoggedIn: (state) => state.isLoggedIn
  },
  
  actions: {
    setUserInfo(userInfo) {
      this.userInfo = { ...this.userInfo, ...userInfo }
    },
    
    setLoginStatus(status) {
      this.isLoggedIn = status
    },
    
    logout() {
      this.userInfo = {
        name: '用户',
        avatar: '',
        email: '',
        phone: ''
      }
      this.isLoggedIn = false
    }
  }
})
