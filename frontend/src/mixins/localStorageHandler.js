
export default {
    methods: {
        isLog() {
            return localStorage.getItem('log') === 'true';
        },
        setLog(value) {
            localStorage.setItem('log', value);
        },
        getLastLogin() {
            return localStorage.getItem('lastLogin')
        },
        setLastLogin(login) {
            localStorage.setItem('lastLogin', login);
        }
    }
}