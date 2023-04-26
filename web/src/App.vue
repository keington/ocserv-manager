<template>
    <div class="container">
        <h3>User management</h3>
        <div class="form-group">
            <label for="username">Username:</label>
            <input type="text" id="username" v-model="username">
            <label for="password">Password:</label>
            <input type="password" id="password" v-model="password">
            <button @click="addUser">Add user</button>
        </div>
        <br>
        <div class="form-group">
            <label for="delete-username">Username:</label>
            <input type="text" id="delete-username" v-model="deleteUsername">
            <button @click="deleteUser">Delete user</button>
        </div>
        <hr>
        <h3>Server management</h3>
        <div class="button-group">
            <button @click="start">Start</button>
            <button @click="stop">Stop</button>
            <button @click="restart">Restart</button>
        </div>
        <br>
        <div class="form-group">
            <input type="text" placeholder="Session ID" v-model="sessionId">
            <button @click="disconnect">Disconnect</button>
        </div>
    </div>
</template>

<script>
import axios from "axios";

export default {
    name: 'App',
    data() {
        return {
            username: '',
            password: '',
            deleteUsername: '',
            sessionId: ''
        };
    },
    methods: {
        sendCommand(command) {
            const url = '/control';
            const params = { command };
            axios.post(url, params)
                .then(response => {
                    console.log(response.data);
                })
                .catch(error => {
                    console.error(error);
                });
        },
        start() {
            this.sendCommand('start');
        },
        stop() {
            this.sendCommand('stop');
        },
        restart() {
            this.sendCommand('restart');
        },
        disconnect() {
            this.sendCommand(`disconnect ${this.sessionId}`);
        },
        addUser() {
            this.sendCommand(`add-user ${this.username} ${this.password}`);
        },
        deleteUser() {
            this.sendCommand(`delete-user ${this.deleteUsername}`);
        }
    }
};
</script>

<style>
.container {
    display: flex;
    flex-direction: column;
    height: 100vh;
    justify-content: center;
    align-items: center;
}

.form-group {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    margin-bottom: 15px;
}

.form-group > * {
    margin-right: 10px;
}

.button-group {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    margin-bottom: 15px;
}

.button-group > * {
    margin-right: 10px;
}
</style>
