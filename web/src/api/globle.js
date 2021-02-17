import service from "@/utils/request";

export const updateRecord = (data) => {
    return service({
        url: "/board/updateRecord",
        method: 'put',
        data: data,
        donNotShowLoading: true
    })
}

export const getRecord = (data) => {
    return service({
        url: "/board/getRecord",
        method: 'get',
        data
    })
}