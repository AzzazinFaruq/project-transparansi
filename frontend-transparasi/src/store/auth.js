import { defineStore } from 'pinia'
import axios from 'axios'
import { useRouter } from 'vue-router'

export const authStore = defineStore('auth', {
  state: () => {
    return{
      user_id:''
    }
  },
  actions:{
    check(router, swal){
      var stat = false;
        axios.get('/api/user')
        .then((res)=>{
          stat = res.data.status;
          this.user_id = res.data.Id
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
