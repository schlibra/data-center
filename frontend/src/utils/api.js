import router from '@/router/index.js'

export function apiBack(res) {
  if (res.status === 200) {
    if (res.data.code === 200) {
      return [true, res.data.data]
    } else if (res.data.code === 401) {
      router.push('/login')
      return [false, res.data.message]
    } else if (res.data.code === 403) {
      router.push('/')
      return [false, res.data.message]
    } else {
      return [false, res.data.message]
    }
  } else {
    return [false, res.statusText]
  }
}