<template>
  <div class="hello">
    <div>
        <input 
            v-model="link"
            type="text"
            placeholder="beanconqueror link"
            @keyup.enter="getShortLink"
        >
        <button @click="getShortLink">
            Generate
        </button>
            <div class="buymeacoffee">
                <div>
                    <div>
                        Do you want to support beanlink's cost of server and domain? Buy me a coffee. 
                    </div>
                    <br>

                    <div>
                        Please do NOT feel obligated to do this.
                        The total cost per YEAR for beanl.ink is about €35 and it is in no way a risk to my personal finances.
                        This is ONLY if you WANT to support me and like the service that beanlink provides.
                    </div>
                </div>
                <a 
                    href="https://www.buymeacoffee.com/beanlink"
                    target="_blank"
                >
                    <img 
                        src="https://cdn.buymeacoffee.com/buttons/v2/arial-yellow.png"
                        alt="Buy Me A Coffee"
                        style="height: 60px !important;width: 217px !important;"
                    >
                </a>
            </div>
    </div>
    <div
        v-if="resultLink"
    >
        <div>
            <p> {{ discordData }} </p>
            {{ resultLink }}
            <button
                @click="copyForDiscord"
            >
                Copy for Discord
            </button>
            <button
                @click="copyToClipboard"
            >
                Copy only link
            </button>
        </div>
    </div>
  </div>

</template>

<script>
import axios from 'axios';

export default {
  name: 'HelloWorld',
    data() {
        return {
            link: '',
            resultLink: null,
            discordData: null,
            publicLink: false
        }
    },
    methods: {
        copyToClipboard() {
            navigator.clipboard.writeText(this.resultLink)
        },
        copyForDiscord() {
            navigator.clipboard.writeText(this.discordData)
        },
        async getShortLink() {
            this.resultLink = null;
            const res = await axios.post('https://beanl.ink/add', {
                link: this.link
            });
            if (res.data.error) {
                window.alert("Your link was not correct");
                return;
            }
            this.resultLink = res.data.link;
            this.discordData = `:beans: ${res.data.name} ${res.data.roaster} <${this.resultLink}>`
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.buymeacoffee {
    position: absolute;
    width:95%;
    bottom:0;
}
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
