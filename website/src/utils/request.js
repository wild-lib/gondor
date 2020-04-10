import axios from 'axios'
import { Message, MessageBox } from 'element-ui'
import store from '@/store'
import { getToken, removeToken } from '@/utils/auth'
import { resetRouter } from '@/router'

// create an axios instance
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API, // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 5000 // request timeout
})

// request interceptor
service.interceptors.request.use(
  config => {
    // do something before request is sent

    if (store.getters.token) {
      // let each request carry token
      // ['X-Token'] is a custom headers key
      // please modify it according to the actual situation
      config.headers['X-Token'] = getToken()
    }
    return config
  },
  error => {
    // do something with request error
    console.log(error) // for debug
    return Promise.reject(error)
  }
)

// response interceptor
service.interceptors.response.use(
  /**
   * If you want to get http information such as headers or status
   * Please return  response => response
   */

  /**
   * Determine the request status by custom code
   * Here is just an example
   * You can also judge the status by HTTP Status Code
   */
  response => {
    const res = response.data
    if (res.code === 200) {
      return res
    }

    // 508: Illegal token; 512: Other clients logged in; 514: Token expired;
    if (res.code === 508 || res.code === 512 || res.code === 514) {
      // to re-login
      MessageBox.confirm(
        '已退出系统, 您可点击“取消”保留在当前页面，或者重新登录。',
        '确认退出',
        {
          confirmButtonText: '重新登录',
          cancelButtonText: '取消',
          type: 'warning'
        }
      ).then(() => {
        store.dispatch('user/resetToken').then(() => {
          location.reload()
        })
      })
    } else { // if the custom code is not 200, it is judged as an error.
      Message({
        message: res.message || '出错了',
        type: 'error',
        duration: 5 * 1000
      })
    }
    return Promise.reject(new Error(res.message || '出错了'))
  },
  error => {
    console.log('err' + error) // for debug
    if (error.response && error.response.status === 401) {
      MessageBox.alert(
        '时效已过，请重新登录！',
        {
          confirmButtonText: '重新登录',
          type: 'warning'
        }
      ).then(() => {
        removeToken()
        resetRouter()
        location.reload()
      })
      error.message = '时效已过，请重新登录！'
    } else {
      Message({
        message: error.message,
        type: 'error',
        duration: 5 * 1000
      })
    }
    return Promise.reject(error)
  }
)

export default service
