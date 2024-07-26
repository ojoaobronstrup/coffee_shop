import { createStore } from "vuex";

const store = createStore({
    state() {
        return {
            jwtToken: "teste"
        }
    }
})

export default store