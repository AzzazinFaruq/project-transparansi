<template>
  <v-container>
   <div class="my-5">
     <v-btn variant="text" @click="router.go(-1)" prepend-icon="mdi-chevron-left">Kembali</v-btn>
   </div>
   <div class="mt-10">
     <v-row>
       <v-col cols="12" md="3">
         <div class="text-center">
           <v-img 
             v-if="!user.foto_profil && !imagePreview" 
             src="@/assets/profile.png" 
             width="80%" 
             height="auto"
             class="mx-auto mb-3"
           ></v-img>
           <v-img 
             v-else-if="imagePreview" 
             :src="imagePreview" 
             width="80%" 
             height="auto"
             class="mx-auto mb-3"
           ></v-img>
           <v-img 
             v-else 
             :src="`${getImageUrl(user.foto_profil)}`" 
             width="80%" 
             height="auto"
             class="mx-auto mb-3"
           ></v-img>

           <v-btn 
             color="primary" 
             size="small"
             @click="dialog = true"
             style="width: 80%"
           >
             Ubah Foto
           </v-btn>
         </div>
       </v-col>
       <v-col cols="12" md="9">
        <v-form @submit.prevent="saveProfile">
         <v-row>
           <v-col cols="12" sm="6">
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
           <v-col cols="12" sm="6">
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
         <div class="mt-2">
            <v-btn
            color="#387144" 
                type="submit" 
                :loading="isSaving"
                style="text-transform: none;">Simpan</v-btn>
            <v-btn color="#BF3232" class="ml-1"  @click="router.go(-1)" style="text-transform: none;">Batal</v-btn>
   </div>     
  </v-form>
       </v-col>
     </v-row>
   </div>

   <v-dialog v-model="dialog" max-width="500px">
      <v-card>
        <v-card-title>Ubah Foto Profil</v-card-title>
        <v-card-text>
          <v-file-input
            v-model="selectedImage"
            accept="image/*"
            label="Pilih Foto"
            @change="previewImage"
            prepend-icon="mdi-camera"
          ></v-file-input>
          
          <v-img 
            v-if="imagePreview"
            :src="imagePreview"
            max-height="200"
            contain
            class="mt-3"
          ></v-img>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="error" @click="closeDialog">Batal</v-btn>
          <v-btn 
            color="primary" 
            @click="confirmImage"
            :disabled="!selectedImage"
          >
            Pilih Foto
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
   
  </v-container>
 </template>
 <script>
 import { useRouter } from 'vue-router';
 import axios from 'axios';
 import { getImageUrl } from '@/config/foto';
 import Swal from 'sweetalert2';
 
 export default {
   setup() {
     const router = useRouter();
     return { router };
   },
   data:()=>({
     getImageUrl,
     user:[],
     jabatan: [],
     selectedImage: null,
     imagePreview: null,
     dialog: false,
     isSaving: false
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
     },
     previewImage() {
       if (this.selectedImage) {
         this.imagePreview = URL.createObjectURL(this.selectedImage);
       } else {
         this.imagePreview = null;
       }
     },
     closeDialog() {
       this.dialog = false;
       this.selectedImage = null;
       this.imagePreview = null;
     },
     confirmImage() {
       this.dialog = false;
     },
     async saveProfile() {
       try {
         this.isSaving = true;
         const formData = new FormData();

         // Append semua field ke FormData
         formData.append('username', this.user.username);
         formData.append('email', this.user.email);
         formData.append('no_hp', this.user.no_hp);
         formData.append('alamat', this.user.alamat);
         if (this.user.jabatan_id) {
           formData.append('jabatan_id', this.user.jabatan_id);
         }
         
         // Append foto jika ada
         if (this.selectedImage) {
           formData.append('foto_profil', this.selectedImage);
         }

         await axios.put(`/api/user/edituser/${this.user.Id}`, formData, {
           headers: {
             'Content-Type': 'multipart/form-data'
           }
         });

         // Refresh data
         await this.Getuser();
         
         // Reset foto states
         this.selectedImage = null;
         this.imagePreview = null;

         await Swal.fire({
          icon: 'success',
          title: 'Berhasil!',
          text: 'Profil berhasil diperbarui',
          showConfirmButton: false,
          timer: 1500
        });
         this.router.go(-1);

       } catch (error) {
         console.error('Error updating profile:', error);
         Swal.fire({
          icon: 'error',
          title: 'Oops...',
          text: 'Gagal memperbarui profil',
          confirmButtonText: 'Tutup',
          confirmButtonColor: '#3085d6'
        });
       } finally {
         this.isSaving = false;
       }
     }
   }
 };
 </script>
 