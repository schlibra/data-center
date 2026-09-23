import axios from 'axios'
import { useTokenStore } from '@/stores/token.js'
import router from '@/router/index.js'

const token = useTokenStore()

async function loginKey(username) {
  try {
    const res = await axios.put("/api/user/login", {
      username
    })
    if (res.data.code === 200) {
      return [true, res.data.data.public_key]
    } else {
      return [false, res.data.message]
    }
  } catch (e) {
    return [false, e]
  }
}
async function loginUser(username, password) {
  try {
    const res = await axios.post("/api/user/login", {
      username,
      password
    })
    if (res.data.code === 200) {
      token.token = res.data.data.token
      return [true, res.data.message]
    } else {
      return [false, res.data.message]
    }
  } catch (e) {
    return [false, e]
  }
}
async function registerKey(username) {
  try {
    const res = await axios.put("/api/user/register", {
      username
    })
    if (res.data.code === 200) {
      return [true, res.data.data.public_key]
    } else {
      return [false, res.data.message]
    }
  } catch (e) {
    return [false, e]
  }
}
async function registerUser(username, password, nickname) {
  try {
    const res = await axios.post("/api/user/register", {
      username,
      password,
      nickname
    })
    if (res.data.code === 200) {
      return [true, res.data.message]
    } else {
      return [false, res.data.message]
    }
  } catch (e) {
    return [false, e]
  }
}
async function getUserInfo() {
  try {
    const res = await axios.get("/api/user", token.config)
    if (res.data.code === 200) {
      return [true, res.data.data.user]
    } else if (res.data.code === 401) {
      router.push("/login")
      return [false, res.data.message]
    } else {
      return [false, res.data.message]
    }
  } catch (e) {
    return [false, e]
  }
}
async function updateUser(nickname) {
  try {
    const res = await axios.put("/api/user", {
      nickname
    }, token.config)
    if (res.data.code === 200) {
      return [true, res.data.message]
    } else {
      return [false, res.data.message]
    }
  } catch (e) {
    return [false, e]
  }
}
async function setUserPassword(password) {
  try {
    const res = await axios.patch("/api/user", {
      password
    }, token.config)
    if (res.data.code === 200) {
      return [true, res.data.message]
    } else {
      return [false, res.data.message]
    }
  } catch (e) {
    return [false, e]
  }
}
async function logoutUser() {
  try {
    const res = await axios.post("/api/user/logout", {}, token.config)
    if (res.data.code === 200) {
      return [true, res.data.message]
    } else {
      return [false, res.data.message]
    }
  } catch (e) {
    return [false, e]
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
  setUserPassword
}