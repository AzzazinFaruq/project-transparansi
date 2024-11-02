<template>
  <v-container>
   <div class="my-5">
     <v-btn variant="text" @click="router.go(-1)" prepend-icon="mdi-chevron-left">Kembali</v-btn>
   </div>
   <div class="mt-10">
     <v-row>
       <v-col cols="12" md="3">
         <div class="">
           <v-img v-if="user.foto_profil == ''" src="@/assets/profile.png" width="80%" height="auto"></v-img>
           <v-img :src="`${getImageUrl(user.foto_profil)}`" width="80%" height="auto"></v-img>
         </div>
       </v-col>
       <v-col cols="12" md="9">
        <v-form>
         <v-row>
           <v-col cols="6">
           <div class="mb-10">
           <label for="" class="label-form">Nama</label>
             <v-text-field v-model="user.username" variant="outlined"></v-text-field>
           </div>
           <div class="mb-10" v-if="user.role_id == 1 || user.role_id == 2">
           <label for="" class="label-form">Jabatan</label>
             <v-select v-model="user.jabatan_id" :items="jabatan" item-title="jabatan" item-value="Id" variant="outlined"></v-select>
           </div>
         <div class="">
           <label for="" class="label-form">No. HP</label>
             <v-text-field v-model="user.no_hp" variant="outlined"></v-text-field>
         </div>
           </v-col>
           <v-col cols="6">
             <div class="mb-10">
           <label for="" class="label-form">Email</label>
           <v-text-field v-model="user.email" variant="outlined"></v-text-field>
         </div>
         <div class="">
           <label for="" class="label-form">Alamat</label>
           <v-textarea v-model="user.alamat" variant="outlined"></v-textarea>
         </div>
           </v-col>
         </v-row>
         </v-form>
         <div class="mt-2">
            <v-btn color="#387144" prepend-icon="mdi-pencil" @click="router.push(`/profile/${user.Id}`)" style="text-transform: none;">Ubah Profil</v-btn>
   </div>     
       </v-col>
     </v-row>
   </div>
   
  </v-container>
 </template>
 <script>
 import { useRouter } from 'vue-router';
 import axios from 'axios';
 import { getImageUrl } from '@/config/foto';
 
 export default {
   setup() {
     const router = useRouter();
     return { router };
   },
   data:()=>({
     getImageUrl,
     user:[],
     jabatan:[]
   }),
   mounted(){
     this.Getuser();
     this.getJabatan();
   },
   methods:{
    getJabatan(){
      axios.get('/api/index-jabatan')
      .then(res=>{
        console.log(res.data.data);
        this.jabatan = res.data.data
      })
    },
     Getuser(){
       axios.get('/api/user')
       .then(res=>{
         this.user = res.data.data
       })
     }
   }
 };
 </script>
 