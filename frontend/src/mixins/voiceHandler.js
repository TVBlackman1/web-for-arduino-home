export default {
    data() {
        return {
            recognition: null,
            funcOnEndListening: null
        }
    },
    methods: {
        initVoiceHandle(onEnd = () => {}) {
            window.SpeechRecognition = window.SpeechRecognition || window.webkitSpeechRecognition;
            this.setSpeechRecognition()
            this.funcOnEndListening = onEnd
        },
        setSpeechRecognition() {
            if (window.SpeechRecognition) {

                /* setup Speech Recognition */
                this.recognition = new window.SpeechRecognition();
                this.recognition.interimResults = true;
                this.recognition.lang = 'ru-RU';
                this.recognition.addEventListener('result', this.transcriptHandler);

                this.recognition.onerror = function(event) {
                    console.log(event.error);

                    /* Revert input and icon CSS if no speech is detected */
                    if(event.error === 'no-speech'){
                        // $voiceTrigger.removeClass('active');
                        // $searchInput.attr("placeholder", "Поиск...");
                    }
                }
            } else {
                // $voiceTrigger.remove();
                console.log("Voice error!");
            }
        },
        transcriptHandler(e) {
            let speechOutput = this.parseTranscript(e)
            // $searchInput.val(speechOutput);
            // console.log(speechOutput)
            //$result.html(speechOutput);
            if (e.results[0].isFinal) {
                // $searchForm.submit();
                console.log(speechOutput)
                this.funcOnEndListening(speechOutput)

            }
        },
        parseTranscript(e) {
            return Array.from(e.results).map(function (result) { return result[0] }).map(function (result) { return result.transcript }).join('')
        },
        listenStart(e) {
            e.preventDefault();
            this.recognition.start();
        },

    }
}