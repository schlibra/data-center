import axios from 'axios'
import { useTokenStore } from '@/stores'
import { apiBack } from '@/utils'

const token = useTokenStore()

async function loginKey(username) {
  try {
    const res = await axios.put('/api/user/login', {
      username,
    })
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function loginUser(username, password) {
  try {
    const res = await axios.post('/api/user/login', {
      username,
      password,
    })
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function registerKey(username) {
  try {
    const res = await axios.put('/api/user/register', {
      username,
    })
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function registerUser(username, password, nickname) {
  try {
    const res = await axios.post('/api/user/register', {
      username,
      password,
      nickname,
    })
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function getUserInfo() {
  try {
    const res = await axios.get('/api/user', token.config)
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function updateUser(nickname) {
  try {
    const res = await axios.put(
      '/api/user',
      {
        nickname,
      },
      token.config,
    )
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function setUserPassword(password) {
  try {
    const res = await axios.patch(
      '/api/user',
      {
        password,
      },
      token.config,
    )
    return apiBack(res)
  } catch (e) {
    return [false, e]
  }
}
async function logoutUser() {
  try {
    const res = await axios.post('/api/user/logout', {}, token.config)
    return [true, res.data.message]
  } catch (e) {
    return [true, e]
  }
}
export {
  loginKey,
  loginUser,
  registerKey,
  registerUser,
  getUserInfo,
  updateUser,
  logoutUser,
  setUserPassword,
}