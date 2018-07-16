<template>
<v-expansion-panel>
  <v-expansion-panel-content  v-for="message in payloadData.slice().reverse()" :key="message.id">
    <div slot="header"><strong>Key:</strong> {{message.key}}</div>
    <div slot="header"><strong>ID:</strong> {{message.id}}</div>
    <v-card>
      <v-card-text><pre class="grey-background">{{message.payload}}</pre></v-card-text>
    </v-card>
  </v-expansion-panel-content>
</v-expansion-panel>
</template>

<script>
export default {
  name: 'HelloWorld',
  created: function () {
    var self = this
    this.ws = new WebSocket('wss://pangea-staging.vumatel.co.za/laurentia/ws?apikey=')
    this.ws.addEventListener('message', function (e) {
      var msg = JSON.parse(e.data)
      self.msg = msg
      self.payloadData.push(msg)
    })
  },
  methods: {
    checkConnection: function () {
      this.ws.send('Ping')
    }
  },
  filters: {
    reverse: function (value) {
      return value
    }
  },
  mounted: function () {
    setInterval(function () {
      this.checkConnection()
    }.bind(this), 10000)
  },
  data () {
    return {
      msg: 'Welcome to Your Vue.js App',
      ws: null,
      payloadData: [],
      chatContent: ''
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
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
.grey-background {
  background-color: lightgrey;
  padding: 10px;
  color: maroon;
}
</style>
