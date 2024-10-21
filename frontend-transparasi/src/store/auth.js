import { defineStore } from 'pinia'
import axios from 'axios'
import { useRouter } from 'vue-router'

export const authStore = defineStore('auth', {
  actions:{
    check(router, swal){
      var stat = false;
        axios.get('/api/user')
        .then((res)=>{
          stat = res.data.status;
        })
        .catch(err=>{
          if (err.response.data.status ==  false) {
            router.push("/login")
            swal({
              toast: "true",
              timer:4000,
              position:"top-end",
              icon: "error",
              title: "Error ",
              text:"Anda Belum Login",
              showConfirmButton :false
              });
          }
        })



    }
  },
})
