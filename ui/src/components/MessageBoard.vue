<template>
    <v-container fluid>
        <v-navigation-drawer app></v-navigation-drawer>
        <v-toolbar app>Message Board</v-toolbar>
        <v-content>
            <v-expansion-panel>
                <v-expansion-panel-content v-for="(item, i) in messageList" :key="i">
                    <div>{{ item.key }}&nbsp;|&nbsp;{{ item.timestamp | moment("dddd, MMMM Do YYYY, h:mm:ss a") }}</div>
                    <v-card>
                        <v-card-text><pre>{{ item.payload }}</pre></v-card-text>
                    </v-card>
                </v-expansion-panel-content>
            </v-expansion-panel>
        </v-content>
        <v-footer app></v-footer>
    </v-container>
</template>

<script>
    // import moment from 'moment'
    export default {
        name: 'MessageBoard',
        props: {
            msg: String
        },
        beforeMount: function () {
            this.loginRedirect()
        },
        mounted: function () {
            let self = this
            this.axios.defaults.headers.common['Authorization'] = 'Bearer ' + this.token
            this.axios.get('https://pangea-staging.vumatel.co.za/laurentia/v1/messages').then((response) => {
                console.log(response.data.messages)
                this.messageList = this.messageList.concat(response.data.messages)
            })
        },
        methods: {
            loginRedirect: function() {
                let token = this.$cookie.get('jwt')
                if (!token){this.$router.push('/login')}
                this.token = token
            }
        },
        data: function()  {
            return {
                token: '',
                messageList:[]
            }

        },
        filter: {

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
