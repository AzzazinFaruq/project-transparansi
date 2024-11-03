<template>
  <div>
    <v-card class="pa-5">
      <div class="d-flex justify-space-between align-center mb-5">
        <div>
          <p class="text-h5 font-weight-bold">Detail Program</p>
          <p class="text-subtitle-1 text-grey">Informasi lengkap program</p>
        </div>
        <v-btn 
          color="error" 
          @click="back()"
          prepend-icon="mdi-arrow-left"
        >
          Kembali
        </v-btn>
      </div>

      <v-divider class="mb-5"></v-divider>

      <v-row>
        <v-col cols="12" md="8">
          <v-form v-if="user.status == 'Menunggu' || user.status == 'Dalam Proses'">
            <!-- Informasi Dasar -->
            <div class="mb-5">
              <label class="label-form mb-2 d-block">Nama Program</label>
              <v-text-field
                v-model="user.nama_program"
                variant="outlined"
                density="compact"
              ></v-text-field>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Institusi</label>
              <v-select
                v-model="user.institusi_id"
                :items="institusi"
                item-title="nama_institusi"
                item-value="id"
                variant="outlined"
                density="compact"
              ></v-select>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Deskripsi</label>
              <v-textarea
                v-model="user.deskripsi"
                variant="outlined"
                density="compact"
              ></v-textarea>
            </div>

            <!-- Detail Anggaran -->
            <h2 class="text-h6 font-weight-bold mb-3">Detail Anggaran</h2>
            <v-divider class="mb-5"></v-divider>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Jenis Anggaran</label>
              <v-select
                v-model="user.jenis_anggaran_id"
                :items="JenisAnggaran"
                item-title="jenis"
                item-value="Id"
                variant="outlined"
                density="compact"
              ></v-select>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Jumlah Anggaran</label>
              <v-text-field
                v-model="user.jumlah_anggaran"
                variant="outlined"
                density="compact"
              ></v-text-field>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Kategori Penggunaan</label>
              <v-select
                v-model="user.kategori_penggunaan_id"
                :items="KategoriPenggunaan"
                item-title="kategori"
                item-value="Id"
                variant="outlined"
                density="compact"
              ></v-select>
            </div>

            <!-- Detail Alamat -->
            <h2 class="text-h6 font-weight-bold mb-3">Detail Alamat</h2>
            <v-divider class="mb-5"></v-divider>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Dusun</label>
              <v-text-field
                v-model="user.dusun"
                variant="outlined"
                density="compact"
              ></v-text-field>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Kabupaten / Kota</label>
              <v-autocomplete
                v-model="user.kabupaten_id"
                :items="kablist"
                item-title="nama_kabupaten"
                item-value="Id"
                variant="outlined"
                density="compact"
              ></v-autocomplete>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Kecamatan</label>
              <v-autocomplete
                v-model="user.kecamatan_id"
                :items="keclist"
                item-title="nama_kecamatan"
                item-value="Id"
                variant="outlined"
                density="compact"
              ></v-autocomplete>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Desa</label>
              <v-autocomplete
                v-model="user.desa_id"
                :items="deslist"
                item-title="nama_desa"
                item-value="Id"
                variant="outlined"
                density="compact"
              ></v-autocomplete>
            </div>
          
            <!-- Detail Dokumentasi -->
            <div v-if="user.status == 'Dalam Proses'">
              <h2 class="text-h6 font-weight-bold mb-3">Detail Dokumentasi</h2>
              <v-divider class="mb-5"></v-divider>

              <div class="foto-section">
                <!-- Foto Before -->
                <div class="mb-5">
                  <label class="label-form mb-2 d-block">Foto Before</label>
                  <div class="preview-container mb-2" v-if="previewBefore || user.foto_before">
                    <img :src="previewBefore || `${getImageUrl(user.foto_before)}`" alt="Preview Before" class="preview-image">
                  </div>
                  <v-file-input
                    v-model="fotoBeforeFile"
                    @change="handleFileUpload($event, 'before')"
                    accept="image/*"
                    variant="outlined"
                    density="compact"
                    prepend-icon="false"
                    prepend-inner-icon="mdi-image"
                  ></v-file-input>
                </div>

                <!-- Foto Progress -->
                <div class="mb-5">
                  <label class="label-form mb-2 d-block">Foto Progress</label>
                  <div class="preview-container mb-2" v-if="previewProgress || user.foto_progress">
                    <img :src="previewProgress || `${getImageUrl(user.foto_progress)}`" alt="Preview Progress" class="preview-image">
                  </div>
                  <v-file-input
                    v-model="fotoProgressFile"
                    @change="handleFileUpload($event, 'progress')"
                    accept="image/*"
                    variant="outlined"
                    density="compact"
                    prepend-icon="false"
                    prepend-inner-icon="mdi-image"
                  ></v-file-input>
                </div>

                <!-- Foto After -->
                <div class="mb-5">
                  <label class="label-form mb-2 d-block">Foto After</label>
                  <div class="preview-container mb-2" v-if="previewAfter || user.foto_after">
                    <img :src="previewAfter || `${getImageUrl(user.foto_after)}`" alt="Preview After" class="preview-image">
                  </div>
                  <v-file-input
                    v-model="fotoAfterFile"
                    @change="handleFileUpload($event, 'after')"
                    accept="image/*"
                    variant="outlined"
                    density="compact"
                    prepend-icon="false"
                    prepend-inner-icon="mdi-image"
                  ></v-file-input>
                </div>
              </div>
            </div>

            <!-- Tombol Aksi -->
          

            <div class="d-flex gap-2 mt-5" v-if="user.status=='Dalam Proses' || user.status == 'Menunggu'">
              <v-btn
                color="#387144"
                style="color: white"
                prepend-icon="mdi-content-save"
                @click=" simpan()"
              >
                Simpan
              </v-btn>
              <v-btn
                color="#BF3232"
                style="color: white"
                @click="back()"
              >
                Kembali
              </v-btn>
            </div>
          </v-form>

          <!-- View Mode untuk status Ditolak atau Selesai -->
          <div v-if="user.status=='Ditolak' || user.status=='Selesai'">
            <!-- Informasi Dasar -->
            <div class="mb-5">
              <label class="label-form mb-2 d-block">Nama Program</label>
              <p class="text-body-1">{{ user.nama_program }}</p>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Institusi</label>
              <p class="text-body-1">{{ getInstitusiName(user.institusi_id) }}</p>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Deskripsi</label>
              <p class="text-body-1">{{ user.deskripsi }}</p>
            </div>

            <!-- Detail Anggaran -->
            <h2 class="text-h6 font-weight-bold mb-3">Detail Anggaran</h2>
            <v-divider class="mb-5"></v-divider>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Jenis Anggaran</label>
              <p class="text-body-1">{{ getJenisAnggaranName(user.jenis_anggaran_id) }}</p>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Jumlah Anggaran</label>
              <p class="text-body-1">Rp {{ user.jumlah_anggaran }}</p>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Kategori Penggunaan</label>
              <p class="text-body-1">{{ getKategoriName(user.kategori_penggunaan_id) }}</p>
            </div>

            <!-- Detail Alamat -->
            <h2 class="text-h6 font-weight-bold mb-3">Detail Alamat</h2>
            <v-divider class="mb-5"></v-divider>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Dusun</label>
              <p class="text-body-1">{{ user.dusun }}</p>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Desa</label>
              <p class="text-body-1">{{ getDesaName(user.desa_id) }}</p>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Kecamatan</label>
              <p class="text-body-1">{{ getKecamatanName(user.kecamatan_id) }}</p>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Kabupaten / Kota</label>
              <p class="text-body-1">{{ getKabupatenName(user.kabupaten_id) }}</p>
            </div>

            <!-- Detail Dokumentasi untuk status Selesai -->
            <div v-if="user.status=='Selesai'">
              <h2 class="text-h6 font-weight-bold mb-3">Detail Dokumentasi</h2>
              <v-divider class="mb-5"></v-divider>

              <div class="foto-section">
                <!-- Foto Before -->
                <div class="mb-5">
                  <label class="label-form mb-2 d-block">Foto Before</label>
                  <div class="preview-container" v-if="user.foto_before">
                    <img :src="getImageUrl(user.foto_before)" alt="Foto Before" class="preview-image">
                  </div>
                </div>

                <!-- Foto Progress -->
                <div class="mb-5">
                  <label class="label-form mb-2 d-block">Foto Progress</label>
                  <div class="preview-container" v-if="user.foto_progress">
                    <img :src="getImageUrl(user.foto_progress)" alt="Foto Progress" class="preview-image">
                  </div>
                </div>

                <!-- Foto After -->
                <div class="mb-5">
                  <label class="label-form mb-2 d-block">Foto After</label>
                  <div class="preview-container" v-if="user.foto_after">
                    <img :src="getImageUrl(user.foto_after)" alt="Foto After" class="preview-image">
                  </div>
                </div>
              </div>
            </div>

            <!-- Status Badge -->
            <div class="mb-5">
              <v-chip
                :color="user.status === 'Selesai' ? '#387144' : '#BF3232'"
                style="color: white"
                class="mt-3"
              >
                {{ user.status }}
              </v-chip>
            </div>

            <!-- Tombol Kembali -->
            <div class="d-flex gap-2 mt-5">
              <v-btn
                color="#BF3232"
                style="color: white"
                @click="back()"
              >
                Kembali
              </v-btn>
            </div>
          </div>
        </v-col>

        <v-col cols="12" md="4">
          <div class="d-flex justify-center">
            <v-img
              src="./logo-dprd-1.png"
              max-width="250"
              contain
              class="mt-5"
            ></v-img>
          </div>
        </v-col>
      </v-row>
    </v-card>
  </div>
</template>

<script>
import axios from 'axios';
import { getImageUrl } from '@/config/foto';
import Swal from 'sweetalert2';

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
      axios.get("/api/institusi")
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
    back(){
      this.$router.go(-1)
    },
    async selesai() {
      try {
        // Konfirmasi selesaikan program
        const result = await Swal.fire({
          title: 'Selesaikan Program?',
          text: "Pastikan semua dokumentasi telah lengkap",
          icon: 'warning',
          showCancelButton: true,
          confirmButtonColor: '#3085d6',
          cancelButtonColor: '#d33',
          confirmButtonText: 'Ya, Selesaikan!',
          cancelButtonText: 'Batal'
        });

        if (result.isConfirmed) {
          // Cek kelengkapan dokumentasi
          if (!this.user.foto_before || !this.user.foto_progress || !this.user.foto_after) {
            Swal.fire({
              icon: 'error',
              title: 'Dokumentasi Belum Lengkap',
              text: 'Harap lengkapi foto before, progress, dan after sebelum menyelesaikan program',
              confirmButtonText: 'OK'
            });
            return;
          }

          // Tampilkan loading
          Swal.fire({
            title: 'Sedang memproses...',
            allowOutsideClick: false,
            didOpen: () => {
              Swal.showLoading();
            }
          });

          // Proses selesaikan program
          await axios.get(`/api/program/selesai/${this.$route.params.id}`);
          
          // Tampilkan sukses
          await Swal.fire({
            icon: 'success',
            title: 'Program Selesai',
            text: 'Program telah berhasil diselesaikan',
            showConfirmButton: false,
            timer: 1500
          });

          this.$router.go(-1);
        }
      } catch (error) {
        console.error('Error:', error);
        Swal.fire({
          icon: 'error',
          title: 'Oops...',
          text: 'Terjadi kesalahan saat menyelesaikan program',
          confirmButtonText: 'Tutup'
        });
      }
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
    },

    getInstitusiName(id) {
      const institusi = this.institusi.find(item => item.id === id)
      return institusi ? institusi.nama_institusi : '-'
    },

    getJenisAnggaranName(id) {
      const jenis = this.JenisAnggaran.find(item => item.Id === id)
      return jenis ? jenis.jenis : '-'
    },

    getKategoriName(id) {
      const kategori = this.KategoriPenggunaan.find(item => item.Id === id)
      return kategori ? kategori.kategori : '-'
    },

    getDesaName(id) {
      const desa = this.deslist.find(item => item.Id === id)
      return desa ? desa.nama_desa : '-'
    },

    getKecamatanName(id) {
      const kecamatan = this.keclist.find(item => item.Id === id)
      return kecamatan ? kecamatan.nama_kecamatan : '-'
    },

    getKabupatenName(id) {
      const kabupaten = this.kablist.find(item => item.Id === id)
      return kabupaten ? kabupaten.nama_kabupaten : '-'
    },

    formatRupiah(value) {
      return new Intl.NumberFormat('id-ID').format(value)
    },

    async getKecamatan(kab_id) {
      try {
        const response = await axios.get(`/api/kecamatan/${kab_id}`)
        this.keclist = response.data.data
      } catch (error) {
        console.error('Error fetching kecamatan:', error)
        Swal.fire({
          icon: 'error',
          title: 'Error',
          text: 'Gagal mengambil data kecamatan'
        })
      }
    },

    async getDesa(kec_id) {
      try {
        const response = await axios.get(`/api/desa/${kec_id}`)
        this.deslist = response.data.data
      } catch (error) {
        console.error('Error fetching desa:', error)
        Swal.fire({
          icon: 'error',
          title: 'Error',
          text: 'Gagal mengambil data desa'
        })
      }
    },

    async listdaerah() {
      try {
        const response = await axios.get('/api/index-kabupaten')
        this.kablist = response.data.data
        
        // Jika ada data program yang dimuat, ambil data kecamatan dan desa
        if (this.user.kabupaten_id) {
          await this.getKecamatan(this.user.kabupaten_id)
          if (this.user.kecamatan_id) {
            await this.getDesa(this.user.kecamatan_id)
          }
        }
      } catch (error) {
        console.error('Error fetching kabupaten:', error)
        Swal.fire({
          icon: 'error',
          title: 'Error',
          text: 'Gagal mengambil data kabupaten'
        })
      }
    },

    async simpan() {
      try {
        // Tampilkan loading
        Swal.fire({
          title: 'Sedang menyimpan...',
          allowOutsideClick: false,
          didOpen: () => {
            Swal.showLoading()
          }
        })

        const formData = new FormData()
        // ... kode formData tetap sama ...

        await axios.put(`/api/program/edit/${this.$route.params.id}`, formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        })

        // Tampilkan sukses
        await Swal.fire({
          icon: 'success',
          title: 'Berhasil!',
          text: 'Program berhasil disimpan',
          timer: 1500,
          showConfirmButton: false
        })

        this.$router.go(-1)
      } catch (error) {
        console.error('Error:', error)
        Swal.fire({
          icon: 'error',
          title: 'Oops...',
          text: 'Terjadi kesalahan saat menyimpan program',
          confirmButtonText: 'Tutup'
        })
      }
    }
  },

  watch: {
    'user.kabupaten_id': {
      handler(newVal) {
        if (newVal) {
          // Reset kecamatan dan desa ketika kabupaten berubah

          this.getKecamatan(newVal)
        }
      }
    },
    'user.kecamatan_id': {
      handler(newVal) {
        if (newVal) {
          // Reset desa ketika kecamatan berubah

          this.getDesa(newVal)
        }
      }
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
.label-form {
  font-size: 14px;
  color: rgba(0, 0, 0, 0.6);
}

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
}

.v-btn {
  text-transform: none !important;
}
</style>

