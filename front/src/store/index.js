import { createStore } from "vuex";

const store = createStore({
    state() {
        return {
            jwtToken: "",
            profileImage: ""
        }
    }
})

export default store