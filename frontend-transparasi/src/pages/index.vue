<template>
  <v-container fluid>

  </v-container>
</template>

<script setup>
import router from '@/router';
import axios from 'axios';
router
 axios.get('/api/user')
        .then((res)=>{
          var stat = res.data.status;
          if (stat ==  true) {
          router.push("/dashboard")
          }
          else{
         router.push("/login")
          }})
</script>
