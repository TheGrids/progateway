import axios from "axios";

const url = process.env.VUE_APP_API;

const userRegistration = (name, surname, email, password) => {
    return axios.post(url + 'auth/signup', {name: name, surname: surname, email: email, password: password}, {headers: { 'Content-Type': 'application/json' }});
}

const userLogin = (email, password) => {
    return axios.post(url + 'auth/signin', {email: email, password: password}, {headers: { 'Content-Type': 'application/json' }});
}

export default {userRegistration, userLogin}