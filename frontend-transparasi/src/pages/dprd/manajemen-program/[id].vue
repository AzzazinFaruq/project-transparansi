<template>
  <v-container>

      <v-row>
        <v-col cols="9">

      <v-form>
      <div class="" v-if="user.status == 'Menunggu' || user.status == 'Dalam Proses'">
        <label class="label-form">Nama Program</label>
        <v-text-field v-model="user.nama_program" variant="outlined"></v-text-field>
        <label class="label-form">Institusi</label>
        <v-select
            v-model="user.institusi_id"
            :items="institusi"
            item-title="nama_institusi"
            item-value="Id"
            variant="outlined"
          ></v-select>
        <label  class="label-form">Deskripsi</label>
        <v-textarea v-model="user.deskripsi" variant="outlined"></v-textarea>

        <h2 class="my-3">Detail Anggaran</h2>

        <label class="label-form">Jenis Anggaran</label>
        <v-select
            v-model="user.jenis_anggaran_id"
            :items="JenisAnggaran"
            item-title="jenis"
            item-value="Id"
            variant="outlined"
          ></v-select>
        <label class="label-form">Jumlah Anggaran</label>
        <v-text-field v-model="user.jumlah_anggaran" variant="outlined"></v-text-field>
        <label class="label-form">kategori Penggunaan</label>
        <v-select
            v-model="user.kategori_penggunaan_id"
            :items="KategoriPenggunaan"
            item-title="kategori"
            item-value="Id"
            variant="outlined"
          ></v-select>

        <h2 class="my-3">Detail Alamat</h2>

        <label class="label-form">Dusun</label>
        <v-text-field v-model="user.dusun" variant="outlined"></v-text-field>
        <label class="label-form">Desa</label>
        <v-autocomplete
          v-model="user.desa_id"
          :items="deslist"
          item-title="nama_desa"
          item-value="Id"
          variant="outlined"
        ></v-autocomplete>
        <label class="label-form">Kecamatan</label>
        <v-autocomplete
          v-model="user.kecamatan_id"
          variant="outlined"
          :items="keclist"
          item-title="nama_kecamatan"
          item-value="Id"

        ></v-autocomplete>
        <label class="label-form">Kabupaten / Kota</label>
        <v-autocomplete
          v-model="user.kabupaten_id"
          variant="outlined"
          :items="kablist"
          item-title="nama_kabupaten"
          item-value="Id"

        ></v-autocomplete>

        <div class="" v-if="user.status == 'Dalam Proses'">
          <h2 class="my-3">Detail Dokumentasi</h2>

          <div class="foto-section">
    <!-- Foto Before -->
    <div class="mb-4">
      <label class="label-form ml-10">Foto Before</label>
      <div class="preview-container ml-10 mb-2" v-if="previewBefore || user.foto_before">
        <img :src="previewBefore || `${getImageUrl(user.foto_before)}`" alt="Preview Before" class="preview-image">
      </div>
      <v-file-input
        prepend-icon="false"
        prepend-inner-icon="mdi-image"
        v-model="fotoBeforeFile"
        @change="handleFileUpload($event, 'before')"
        accept="image/*"
        variant="outlined"
      ></v-file-input>
    </div>

    <!-- Foto Progress -->
    <div class="mb-4">
      <label class="label-form ml-10">Foto Progress</label>
      <div class="preview-container ml-10 mb-2" v-if="previewProgress || user.foto_progress">
        <img :src="previewProgress || `${getImageUrl(user.foto_before)}`" alt="Preview Progress" class="preview-image">
      </div>
      <v-file-input
        prepend-icon="false"
        prepend-inner-icon="mdi-image"
        v-model="fotoProgressFile"
        @change="handleFileUpload($event, 'progress')"
        accept="image/*"
        variant="outlined"
      ></v-file-input>
    </div>

    <!-- Foto After -->
    <div class="mb-4">
      <label class="label-form ml-10">Foto After</label>
      <div class="preview-container ml-10 mb-2" v-if="previewAfter || user.foto_after">
        <img :src="previewAfter || `${getImageUrl(user.foto_after)}`" alt="Preview After" class="preview-image">
      </div>
      <v-file-input
        prepend-icon="false"
        prepend-inner-icon="mdi-image"
        v-model="fotoAfterFile"
        @change="handleFileUpload($event, 'after')"
        accept="image/*"
        variant="outlined"
      ></v-file-input>
    </div>
  </div>
        </div>

        <div class="d-flex justify-start" v-if="user.status=='Dalam Proses' || user.status=='Menunggu'">
          <v-btn
            variant="tonal"
            style="background-color:#387144;color: white;text-transform: none;"
            @click="isFormComplete ? selesai() : simpan()"
          >
            {{ isFormComplete ? 'Selesai' : 'Simpan' }}
          </v-btn>
            <v-btn variant="tonal" class="ml-2" style="background-color:#BF3232;color: white;text-transform: none;" @click="back()">
              Kembali
            </v-btn>
        </div>
      </div>
      </v-form>



    <div class="" v-if="user.status=='Ditolak'">
      <label class="label-form">Nama Program</label>
      <p>{{ user.nama_program }}</p>
      <label class="label-form">Institusi</label>
      <p>{{ user.Institusi.nama_institusi }}</p>
      <label class="label-form">Deskripsi</label>
      <p>{{ user.deskripsi }}</p>
      <label class="label-form">Deskripsi</label>
      <p>{{ user.deskripsi }}</p>

      <h2 class="my-3">Detail Anggaran</h2>
      <VDivider/>

      <label class="label-form">Jenis Anggaran</label>
      <p>{{ user.JenisAnggaran.jenis }}</p>
      <label class="label-form">Jumlah Anggaran</label>
      <p>{{ user.jumlah_anggaran }}</p>
      <label class="label-form">Kategori Penggunaan</label>
      <p>{{ user.KategoriPenggunaan.kategori}}</p>

      <h2 class="my-3">Detail Alamat</h2>
      <VDivider/>

      <label class="label-form">Dusun</label>
      <p>{{ user.dusun }}</p>
      <label class="label-form">Desa</label>
      <p>{{ user.Desa.nama_desa }}</p>
      <label class="label-form">Kecamatan</label>
      <p>{{ user.Kecamatan.nama_kecamatan }}</p>
      <label class="label-form">Kabupaten</label>
      <p>{{ user.Kabupaten.nama_kabupaten  }}</p>

      </div>
        </v-col>
        <v-col cols="3">
          <div class="d-flex justify-center" style="width: 100%;">
          <img src="./logo-dprd-1.png" alt="" style="width: 80%;">
        </div>
        </v-col>
      </v-row>


  </v-container>
</template>
<script>
import axios from 'axios';
import { getImageUrl } from '@/config/foto';
export default{
  data() {
    return {
      getImageUrl,
      institusi:[],
      kablist:[],
      keclist:[],
      deslist:[],
      JenisAnggaran:[],
      KategoriPenggunaan:[],
      user:{
        nama_program:'',
        deskripsi: '',
        jenis_anggaran_id:'',
        jumlah_anggaran:'',
        kategori_penggunaan_id:'',
        institusi_id:'',
        dusun:'',
        desa_id:'',
        kecamatan_id:'',
        kabupaten_id:'',
        foto_before:'',
        foto_progress:'',
        foto_after:''
      },
      fotoBeforeFile: null,
      fotoProgressFile: null,
      fotoAfterFile: null,
      previewBefore: null,
      previewProgress: null,
      previewAfter: null
    }
  },
  mounted() {
    this.getProgram();
    this.listJenisAnggaran();
    this.listKategoriPenggunaan();
    this.listinstitusi();
    this.listdaerah()
  },
  methods: {
    listdaerah(){
      axios.all([
    axios.get('/api/index-kabupaten'),  // Endpoint untuk kabupaten
    axios.get('/api/index-kecamatan'),  // Endpoint untuk kecamatan
    axios.get('/api/index-desa')        // Endpoint untuk desa
  ])
  .then(axios.spread((kabupatenRes, kecamatanRes, desaRes) => {
    this.kablist = kabupatenRes.data.data;
    this.keclist = kecamatanRes.data.data;
    this.deslist = desaRes.data.data;

  }))
  .catch(error => {
    console.error("Error fetching data:", error);
  });
    },
    listinstitusi(){
      axios.get("/api/index-institusi")
      .then(res=>{
        this.institusi = res.data.data
      })
    },
    listJenisAnggaran(){
      axios.get("/api/index-jenis-anggaran")
      .then(res=>{
        this.JenisAnggaran = res.data.data
      })
    },
    listKategoriPenggunaan(){
      axios.get("/api/index-kategori-penggunaan")
      .then(res=>{
        this.KategoriPenggunaan = res.data.data
      })
    },
    getProgram(){
      axios.get(`/api/program/${this.$route.params.id}`)
      .then(res=>{
        console.log(res.data.data)
        this.user = res.data.data
      })
    },
    accept(){
    axios.get(`/api/program/accept/${this.$route.params.id}`)
    .then(res=>{
      this.$router.go(0)
    })
  },
  rejectprogram(){
    axios.get(`/api/program/reject/${this.$route.params.id}`)
    .then(this.$router.go(-1))
  },
  back(){
    this.$router.go(-1)
  },
  selesai(){
    axios.get(`/api/program/selesai/${this.$route.params.id}`)
    this.$router.go(-1)
  },
  handleFileUpload(event, type) {
      const file = event.target.files[0]
      if (file) {
        // Update file untuk upload
        this[`foto${type.charAt(0).toUpperCase() + type.slice(1)}File`] = file

        // Buat preview untuk file baru
        const reader = new FileReader()
        reader.onload = (e) => {
          this[`preview${type.charAt(0).toUpperCase() + type.slice(1)}`] = e.target.result
        }
        reader.readAsDataURL(file)
      } else {
        // Reset preview jika file dihapus
        this[`preview${type.charAt(0).toUpperCase() + type.slice(1)}`] = null
      }
    },

  simpan(){
  // Buat objek FormData baru
  const formData = new FormData()

  // Append semua field dari user object ke FormData
  formData.append('nama_program', this.user.nama_program)
  formData.append('deskripsi', this.user.deskripsi)
  formData.append('jenis_anggaran_id', this.user.jenis_anggaran_id)
  formData.append('jumlah_anggaran', this.user.jumlah_anggaran)
  formData.append('kategori_penggunaan_id', this.user.kategori_penggunaan_id)
  formData.append('institusi_id', this.user.institusi_id)
  formData.append('dusun', this.user.dusun)
  formData.append('desa_id', this.user.desa_id)
  formData.append('kecamatan_id', this.user.kecamatan_id)
  formData.append('kabupaten_id', this.user.kabupaten_id)

  Object.keys(this.user).forEach(key => {
        if (key !== 'foto_before' && key !== 'foto_progress' && key !== 'foto_after') {
          formData.append(key, this.user[key])
        }
      })

      // Append file foto jika ada perubahan
      if (this.fotoBeforeFile) {
        formData.append('foto_before', this.fotoBeforeFile)
      }
      if (this.fotoProgressFile) {
        formData.append('foto_progress', this.fotoProgressFile)
      }
      if (this.fotoAfterFile) {
        formData.append('foto_after', this.fotoAfterFile)
      }

  // Kirim dengan axios menggunakan FormData
  axios.put(`/api/program/edit/${this.$route.params.id}`, formData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
  .then(res => {
    this.$router.go(-1)
  })
}
  },

  computed: {
    isFormComplete() {
      // Cek apakah semua field required sudah terisi
      const requiredFields = {
        nama_program: this.user.nama_program,
        deskripsi: this.user.deskripsi,
        jenis_anggaran_id: this.user.jenis_anggaran_id,
        jumlah_anggaran: this.user.jumlah_anggaran,
        kategori_penggunaan_id: this.user.kategori_penggunaan_id,
        institusi_id: this.user.institusi_id,
        dusun: this.user.dusun,
        desa_id: this.user.desa_id,
        kecamatan_id: this.user.kecamatan_id,
        kabupaten_id: this.user.kabupaten_id
      }

      // Cek status dan tambahan field foto jika status "Dalam Proses"

        const fotoFields = {
          foto_before: this.user.foto_before || this.fotoBeforeFile,
          foto_progress: this.user.foto_progress || this.fotoProgressFile,
          foto_after: this.user.foto_after || this.fotoAfterFile
        }
        Object.assign(requiredFields, fotoFields)


      // Cek apakah semua field terisi (tidak kosong)
      return Object.values(requiredFields).every(field => {
        if (typeof field === 'string') return field.trim() !== ''
        return field !== null && field !== undefined;
      })
    }
  },


}
</script>


<style scoped>
.preview-container {
  width: 200px;
  height: 200px;
  border: 1px solid #ddd;
  border-radius: 4px;
  overflow: hidden;
}

.preview-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.foto-section {
  margin: 20px 0;
  margin-left: -39px;
}


</style>

