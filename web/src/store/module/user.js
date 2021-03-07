import {login} from '@/api/user'
import {jsonInBlacklist} from '@/api/jwt'
import router from '@/router/index'

export const user = {
    namespaced: true,
    state: {
        userInfo: {
            uuid: "",
            username: "",
            authority: "",
        },
        token: "",
        firstDay: ""
    },
    mutations: {
        setUserInfo(state, userInfo) {
            // 这里的 `state` 对象是模块的局部状态
            state.userInfo = userInfo
        },
        setToken(state, token) {
            // 这里的 `state` 对象是模块的局部状态
            state.token = token
        },
        setFirstDay(state, firstDay) {
            state.firstDay = firstDay
        },
        LoginOut(state) {
            state.userInfo = {}
            state.token = ""
            sessionStorage.clear()
            router.push({name: 'login', replace: true})
            window.location.reload()
        },
        ResetUserInfo(state, userInfo = {}) {
            state.userInfo = {
                ...state.userInfo,
                ...userInfo
            }
        }
    },
    actions: {
        async LoginIn({commit, dispatch, rootGetters, getters}, loginInfo) {
            const res = await login(loginInfo)
            if (res.code == 0) {
                commit('setUserInfo', res.data.user)
                commit('setToken', res.data.token)
                commit('setFirstDay', res.data.first_day)
                await dispatch('router/SetAsyncRouter', {}, {root: true})
                const asyncRouters = rootGetters['router/asyncRouters']
                router.addRoutes(asyncRouters)
                // const redirect = router.history.current.query.redirect
                // console.log(redirect)
                // if (redirect) {
                //     router.push({ path: redirect })
                // } else {
                router.push({name: getters["userInfo"].authority.defaultRouter})
                // }
                return true
            }
        },
        async LoginOut({commit}) {
            commit("LoginOut")
            jsonInBlacklist()
        }
    },
    getters: {
        userInfo(state) {
            return state.userInfo
        },
        token(state) {
            return state.token
        },

    }
}