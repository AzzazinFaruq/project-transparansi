import { defineStore } from 'pinia'
import axios from 'axios'
import { useRouter } from 'vue-router'

export const authStore = defineStore('auth', {
  actions:{
    check(router){
      var stat = false;
        axios.get('/api/user')
        .then((res)=>{
          console.log(res.data.status)
          stat = res.data.status
          if (stat ==  true) {
            router.push("/dashboard")
          }
          else{
            router.push("/login")
          }
        })



    }
  },
})
