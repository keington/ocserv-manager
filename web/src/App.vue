<template>
  <div class="main">
    <button @click="startOpenConnect">Start OpenConnect</button>
    <button @click="stopOpenConnect">Stop OpenConnect</button>
    <div class="users">
      <h3>User List:</h3>
      <ul>
        <li v-for="(user, index) in userList" :key="index">{{ user }}</li>
      </ul>
      <form>
        <h3>Add User:</h3>
        <div class="form-group">
          <label>Username:</label>
          <input type="text" v-model="newUser.username">
        </div>
        <div class="form-group">
          <label>Password:</label>
          <input type="password" v-model="newUser.password">
        </div>
        <button @click.prevent="addUser">Add User</button>
      </form>
      <form>
        <h3>Delete User:</h3>
        <div class="form-group">
          <label>Username:</label>
          <input type="text" v-model="deleteUser.username">
        </div>
        <button @click.prevent="deleteUser">Delete User</button>
      </form>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "App",
  data() {
    return {
      userList: [],
      newUser: {
        username: "",
        password: ""
      },
      deleteUser: {
        username: ""
      }
    };
  },
  mounted() {
    this.getUserList();
  },
  methods: {
    getUserList() {
      axios.get("/api/user/list").then(response => {
        this.userList = response.data.message.split("\n");
      });
    },
    startOpenConnect() {
      axios.post("/api/start").then(() => {
        alert("OpenConnect started");
      });
    },
    stopOpenConnect() {
      axios.post("/api/stop").then(() => {
        alert("OpenConnect stopped");
      });
    },
    addUser() {
      axios
        .post(
          `/api/user/add/${this.newUser.username}/${this.newUser.password}`
        )
        .then(() => {
          alert("User added");
          this.getUserList();
          this.newUser.username = "";
          this.newUser.password = "";
        });
    },
    deleteUser() {
      axios
        .delete(`/api/user/delete/${this.deleteUser.username}`)
        .then(() => {
          alert("User deleted");
          this.getUserList();
          this.deleteUser.username = "";
        });
    }
  }
};
</script>

<style scoped>
.main {
  padding: 20px;
}
.form-group {
  margin-bottom: 10px;
}
.users {
  margin-top: 20px;
}
ul {
  margin-top: 0;
  margin-bottom: 10px;
  list-style: none;
}
li {
  margin-bottom: 5px;
}
</style>