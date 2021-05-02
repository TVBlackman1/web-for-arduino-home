
export default {
    methods: {
        isLog() {
            return localStorage.getItem('log') ?? false;
        },
        setLog(log) {
            localStorage.setItem('log', log);
        }
    }
}