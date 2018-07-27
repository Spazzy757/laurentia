<template>
    <v-content>
        <v-container fluid fill-height>
            <v-layout align-center justify-center>
                <v-flex xs12 sm8 md4>
                    <v-card class="elevation-12">
                        <v-toolbar dark color="dark">
                            <v-toolbar-title>Login</v-toolbar-title>
                        </v-toolbar>
                        <v-card-text>
                            <v-form>
                                <v-text-field
                                        prepend-icon="person"
                                        name="login"
                                        label="Login"
                                        type="text"
                                        v-model="loginData.username"
                                ></v-text-field>
                                <v-text-field
                                        prepend-icon="lock"
                                        name="password"
                                        label="Password"
                                        id="password"
                                        type="password"
                                        v-model="loginData.password"
                                ></v-text-field>
                            </v-form>
                        </v-card-text>
                        <v-card-actions>
                            <v-spacer></v-spacer>
                            <v-btn color="dark" @click="login">Login</v-btn>
                        </v-card-actions>
                    </v-card>
                </v-flex>
            </v-layout>
        </v-container>
    </v-content>
</template>

<script>
    export default {
        name: 'Login',
        props: {
            msg: String
        },
        methods: {
            login: function () {
                this.axios.post('https://pangea-staging.vumatel.co.za/laurentia/login', this.loginData).then((response) => {
                    console.log(response.data)
                    this.axios.defaults.headers.common['Authorization'] = 'Bearer ' + response.data.token
                    this.$cookie.set('jwt', response.data.token, response.data.expire)
                    this.$router.push('/')
                })

            }
        },
        data: function()  {
            return {
                loginData: {
                    username: '',
                    password: ''
                }
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
