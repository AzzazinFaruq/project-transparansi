import { defineStore } from 'pinia'
import axios from 'axios'
import { useRouter } from 'vue-router'

export const authStore = defineStore('auth', {
  actions:{
    check(router, swal){
      var stat = false;
        axios.get('/api/user')
        .then((res)=>{
          console.log(res.data.status)
          stat = res.data.status;
          var name = res.data.username;
          if (stat ==  false) {
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
          else{

          }
        })



    }
  },
})
