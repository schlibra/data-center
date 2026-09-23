import { JSEncrypt } from 'jsencrypt'

export default function encrypt(data, pubKey) {
  const e = new JSEncrypt()
  e.setPublicKey(pubKey)

  const ed = e.encrypt(data)
  if (!ed) {
    return [false, "加密失败"]
  }
  return [true, ed]
}