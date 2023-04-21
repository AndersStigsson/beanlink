<template>
  <div class="hello">
    <div>
        <input 
            v-model="link"
            type="text"
            placeholder="beanconqueror link"
            @keyup.enter="getShortLink"
        >
    </div>
    <div
        v-if="resultLink"
    >
        {{ resultLink }}
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
            resultLink: null
        }
    },
    methods: {
        async getShortLink() {
            this.resultLink = null;
            const res = await axios.post('https://backend.beanl.ink/add', {
                link: this.link
            });
            if (res.data.error) {
                window.alert("Your link was not correct");
                return;
            }
            this.resultLink = res.data.link;
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
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
