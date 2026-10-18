function calcTime(time) {
    const baseTimeStamp = 946656000000
    const baseTime = new Date(baseTimeStamp)
    const diffTime = new Date(baseTimeStamp + time * 1000)
    const year = diffTime.getFullYear() - baseTime.getFullYear()
    const month = diffTime.getMonth() - baseTime.getMonth()
    const day = diffTime.getDate() - baseTime.getDate()
    const hour = diffTime.getHours() - baseTime.getHours()
    const minute = diffTime.getMinutes() - baseTime.getMinutes()
    const second = diffTime.getSeconds() - baseTime.getSeconds()
    return `${year ? year + '年' : ''}${month ? month + '月' : ''}${day ? day + '日' : ''}${hour ? hour + '时' : ''}${minute ? minute + '分' : ''}${second ? second + '秒' : ''}`
}

export {
    calcTime
}