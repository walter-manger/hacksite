<template>
  <div>
    <LoginHeader />
    <div class='container'>
      <div class='card'>
        <section class='header'>
          <h4> Create Your Account </h4>
        </section>
        <section class='form'>
          <div class='input-container'>
            <span class='label'> Username </span>
            <input placeholder='Username' v-model="account.Username">
          </div>
          <div class='input-container'>
            <span class='label'> Email </span>
            <input placeholder='someone@email.com' v-model="account.Email">
          </div>
          <div class='input-container'>
            <span class='label'> Password </span>
            <input placeholder='Password' type='password' v-model="account.Password">
          </div>
          <div class='input-container'>
            <span class='label'> Confirm your password </span>
            <input placeholder='Confirm Password' type='password' v-model="account.ConfirmPassword">
          </div>
        </section>
        <section class='message-container'>
          <span class='message'> All fields are required. </span>
          <br>
          <span class='message' v-bind:class="{success: success}"> {{ message }} </span>
        </section>
        <div class='menu-bar'>
          <button class='menu-button' @click="Create" :disabled="!valid"> Create </button>
          <router-link to='/'> Cancel </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import router from '@/router';
import { CreateAccount } from '@/database';
import LoginHeader from '@/components/elements/LoginHeader';

export default {
  components: {
    'LoginHeader': LoginHeader
  },
  data () {
    return {
      account: {
        Username: '',
        Email: '',
        Password: '',
        ConfirmPassword: ''
      },
      valid: false,
      message: undefined,
      success: false
    };
  },
  updated () {
    // TODO: Username validation (no spaces)
    // TODO: Password validation (length, characters)
    // TODO: Shouldn't send "ConfirmPassword" to the API since it isn't needed
    this.valid =
      this.account.Username !== '' &&
      this.account.Email !== '' &&
      this.account.Password !== '' &&
      this.account.ConfirmPassword !== '' &&
      this.ConfirmPasswords();
  },
  methods: {
    Create: function () {
      CreateAccount(this.account)
        .then((response) => {
          if (response.status === 201) {
            this.message = 'Account successfully created. Redirecting to login page...';
            this.success = true;
            this.Redirect();
          } else {
            this.message = 'Something went wrong, please try again';
          }
        }, (error) => {
          console.log('Error: ', error);
        });
    },
    ConfirmPasswords: function () {
      return this.account.Password === this.account.ConfirmPassword;
    },
    Redirect: function () {
      setTimeout(() => {
        router.push('/login');
      }, 3000);
    }
  }
};
</script>

<style scoped>
.container {
  margin: 50px;
}
.card {
  padding: 30px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
}
input {
  display: block;
  font-size: 16px;
  padding: 5px 10px;
  width: 500px;
  border: none;
  font-weight: 100;
  outline: none;
  color: #325778;
  padding-bottom: 10px;
  border-bottom: 1px solid #eee;
  margin-top: 10px;
  transition: all 0.5s linear;
}
.input-container {
  margin-top: 25px;
}
.label {
  font-size: 14px;
  margin-left: 5px;
  font-weight: 500;
}
input:focus {
  border-bottom: 1px solid #325778;
}
a, a:visited {
  color: #325778;
}
a:hover {
  color: #1B3F60;
}
button, a {
  border-radius: 0;
  border: none;
}
button:hover {
  cursor: pointer;
}
.menu-bar {
  display: flex;
  margin-top: 25px;
  align-items: center;
  margin-left: 10px;
}
.menu-button {
  background-color: #529A7F;
  padding: 10px;
  line-height: 16px;
  font-size: 16px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
  color: #fff;
  margin-right: 10px;
}
.menu-button:disabled,
.menu-button:disabled:hover {
  background-color: #919191;
  box-shadow: none;
  cursor: not-allowed;
}
.menu-button:hover {
  background-color: #176548;
}
.message-container {
  margin-top: 25px;
  margin-left: 25px;
}
.message {
  display: inline-block;
  font-style: italic;
  font-size: 14px;
  color: #ff4949;
}
.success {
  color: #4e9155;
}
</style>
