import axios from 'axios'
import { useTokenStore } from '@/stores'
import { apiBack, apiError } from '@/utils'

const token = useTokenStore()

async function loginKey(username) {
  try {
    return apiBack(await axios.put('/api/user/login', {
      username,
    }))
  } catch (e) {
    return apiError(e)
  }
}
async function loginUser(username, password) {
  try {
    return apiBack(await axios.post('/api/user/login', {
      username,
      password,
    }))
  } catch (e) {
    return apiError(e)
  }
}
async function registerKey(username) {
  try {
    return apiBack(await axios.put('/api/user/register', {
      username,
    }))
  } catch (e) {
    return apiError(e)
  }
}
async function registerUser(username, password, nickname) {
  try {
    return apiBack(await axios.post('/api/user/register', {
      username,
      password,
      nickname,
    }))
  } catch (e) {
    return apiError(e)
  }
}
async function getUserInfo() {
  try {
    return apiBack(await axios.get('/api/user', token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function updateUser(nickname) {
  try {
    return apiBack(await axios.put(
      '/api/user',
      {
        nickname,
      },
      token.config,
    ))
  } catch (e) {
    return apiError(e)
  }
}
async function setUserPassword(password) {
  try {
    return apiBack(await axios.patch(
      '/api/user',
      {
        password,
      },
      token.config,
    ))
  } catch (e) {
    return apiError(e)
  }
}
async function logoutUser() {
  try {
    return apiBack(await axios.post('/api/user/logout', {}, token.config))
  } catch (e) {
    return apiError(e)
  }
}
async function getUserApiKey() {
  try {
    return apiBack(await axios.post('/api/user/api', {}, token.config))
  } catch (e) {
    return apiError(e)
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
  getUserApiKey,
}