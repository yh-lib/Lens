// 对象转列表
// {k1 : v1,k2 : v2}
// [{key:k1,value:v1},{key:k2,value:v2}]
export const obj2list = (obj) => {
    let list = []
    if (obj == (null || undefined)) {
        return list
    }
    for (let [key, value] of Object.entries(obj)) {
        const newObj = {
            key: key,
            value: value
        }
        list.push(newObj)
    }
    return list
}

// 列表转对象
// [{key:k1,value:v1},{key:k2,value:v2}]
// {k1 : v1,k2 : v2}
export const list2obj = (list) => {
    const obj = {}
    list.forEach(({ key, value }) => {
        if (!key) return
        obj[key] = value
    })
    return obj
}