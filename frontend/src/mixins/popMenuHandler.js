import popMenuOpener from "../emitters/popMenuOpener";

export default {
    methods: {
        __setContent(content) {
            popMenuOpener.$emit('set-content-pop-menu', content)
        },
        __toggle() {
            popMenuOpener.$emit('toggle-pop-menu')
        }
    }
}