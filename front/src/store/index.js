import { createStore } from "vuex";

const store = createStore({
    state() {
        return {
            jwtToken: "",
            profileImage: "",
            username: ""
        }
    }
})

export default store