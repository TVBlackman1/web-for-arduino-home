import mainEmitter from "../emitters/mainEmitter";

export default {
    methods: {
        __setContent(content) {
            mainEmitter.$emit('set-content-pop-menu', content)
        },
        __toggle() {
            mainEmitter.$emit('toggle-pop-menu')
        }
    }
}