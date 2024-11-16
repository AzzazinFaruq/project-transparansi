<template>
  <div>
    <div class="">
      <v-btn variant="text" prepend-icon="mdi-chevron-left" @click="this.$router.go(-1)">Kembali</v-btn>
    </div>
      <v-row>
        <v-col cols="12" md="8">
          <v-card elevation="4" class="pa-4 ma-4">
          <v-form >
            <!-- Informasi Dasar -->
            <div class="mb-5">
              <label class="label-form mb-2 d-block">Nama Program / Proyek</label>
              <v-text-field
                v-model="user.nama_program"
                variant="outlined"
                density="compact"
              ></v-text-field>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Nama Aspirator</label>
              <v-select
                v-model="user.aspirator_id"
                :items="aspirator"
                item-title="nama_aspirator"
                item-value="Id"
                variant="outlined"
                density="compact"
              ></v-select>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Nama Dinas Verifikator</label>
              <v-select
                v-model="user.dinas_verifikator_id"
                :items="verifikator"
                item-title="nama_dinas_verifikator"
                item-value="Id"
                variant="outlined"
                density="compact"
              ></v-select>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Institusi</label>
              <v-text-field
                v-model="user.nama_institusi"
                variant="outlined"
                density="compact"
              ></v-text-field>
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
                type="number"
                min="0"
                @input="formatJumlahAnggaran"
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

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Dusun</label>
              <v-text-field
                v-model="user.dusun"
                variant="outlined"
                density="compact"
              ></v-text-field>
            </div>

            <!-- Tambahkan setelah Detail Alamat dan sebelum form dusun -->
            <div class="mb-5">
              <label class="label-form mb-2 d-block">Pilih Lokasi di Peta</label>
              <div class="search-container mb-2">
                <v-text-field
                  v-model="searchQuery"
                  label="Cari lokasi..."
                  variant="outlined"
                  density="compact"
                  @keyup.enter="handleSearch"
                  append-inner-icon="mdi-magnify"
                  @click:append-inner="handleSearch"
                ></v-text-field>
              </div>
              <div id="map" style="height: 400px;" class="mb-3"></div>
            </div>

            <!-- Detail Dokumentasi -->
            <div >
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


            <div class="d-flex justify-space-between gap-2 mt-5" >
              <v-btn
                color="#387144"
                style="color: white"
                prepend-icon="mdi-content-save"
                @click=" savebutton()"
              >
                Simpan
              </v-btn>
              <div class="">
                <v-btn
                color="#BF3232"
                style="color: white"
                prepend-icon="mdi-content-save"
                @click=" draft()"
              >
                Draft
              </v-btn>
              <v-btn
                color="#387144"
                style="color: white"
                @click="publish()"
              >
                Publish
              </v-btn>
            </div>
            </div>
          </v-form>
        </v-card>
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
  </div>
</template>

<script>
import axios from 'axios';
import { getImageUrl } from '@/config/foto';
import Swal from 'sweetalert2';
import router from '@/router';
import 'leaflet/dist/leaflet.css'
import L from 'leaflet'
import { OpenStreetMapProvider } from 'leaflet-geosearch';

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
        aspirator_id:'',
        dinas_verifikator_id:'',
        deskripsi: '',
        jenis_anggaran_id:'',
        jumlah_anggaran:'',
        kategori_penggunaan_id:'',
        nama_institusi:'',
        dusun:'',
        desa_id:'',
        kecamatan_id:'',
        kabupaten_id:'',
        foto_before:'',
        foto_progress:'',
        foto_after:'',
        latitude: '',
        longitude: '',
      },
      fotoBeforeFile: null,
      fotoProgressFile: null,
      fotoAfterFile: null,
      previewBefore: null,
      previewProgress: null,
      previewAfter: null,
      aspirator:[],
      verifikator:[],
      map: null,
      marker: null,
      searchQuery: '',
      searchProvider: null,
    }
  },
  mounted() {
    this.getProgram();
    this.listJenisAnggaran();
    this.listKategoriPenggunaan();
    this.listdaerah()
    this.aspiratorList()
    this.verifikatorList()
  },
  methods: {
    verifikatorList(){
      axios.get("/api/index-dinas-verifikator")
      .then(res=>{
        this.verifikator = res.data.data
      })
    },
    aspiratorList(){
      axios.get("/api/index-aspirator")
      .then(res=>{
        this.aspirator = res.data.data
      })
    },
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
    async getProgram() {
      try {
        const response = await axios.get(`/api/program/${this.$route.params.id}`);
        this.user = response.data.data;

        // Inisialisasi peta setelah data terload
        this.$nextTick(() => {
          this.initMap();
        });
      } catch (error) {
        console.error('Error fetching program:', error);
      }
    },
    back(){
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
    async savebutton() {
      await this.saveProgram('/api/program/edit/');
    },
    async draft() {
      await this.saveProgram('/api/program/draft/');
    },
    async publish() {
      await this.saveProgram('/api/program/publish/');
    },
    async saveProgram(apiEndpoint) {
      try {
        // Validasi koordinat jika diisi
        if (this.user.latitude || this.user.longitude) {
          const lat = parseFloat(this.user.latitude);
          const lng = parseFloat(this.user.longitude);

          if (isNaN(lat) || isNaN(lng) ||
              lat < -90 || lat > 90 ||
              lng < -180 || lng > 180) {
            await Swal.fire({
              icon: 'error',
              title: 'Koordinat Tidak Valid',
              text: 'Silakan pilih lokasi yang valid pada peta'
            });
            return;
          }
        }

        // Tampilkan loading
        Swal.fire({
          title: 'Sedang menyimpan...',
          allowOutsideClick: false,
          didOpen: () => {
            Swal.showLoading();
          }
        });

        const formData = new FormData();
        // Menambahkan data ke formData
        formData.append('nama_program', this.user.nama_program);
        formData.append('aspirator_id', this.user.aspirator_id);
        formData.append('dinas_verifikator_id', this.user.dinas_verifikator_id);
        formData.append('deskripsi', this.user.deskripsi);
        formData.append('jenis_anggaran_id', this.user.jenis_anggaran_id);
        formData.append('jumlah_anggaran', this.user.jumlah_anggaran);
        formData.append('kategori_penggunaan_id', this.user.kategori_penggunaan_id);
        formData.append('nama_institusi', this.user.nama_institusi);
        formData.append('dusun', this.user.dusun);
        formData.append('desa_id', this.user.desa_id);
        formData.append('kecamatan_id', this.user.kecamatan_id);
        formData.append('kabupaten_id', this.user.kabupaten_id);

        // Tambahkan koordinat ke formData
        formData.append('latitude', this.user.latitude);
        formData.append('longitude', this.user.longitude);

        // Menambahkan file jika ada
        if (this.fotoBeforeFile) {
          formData.append('foto_before', this.fotoBeforeFile);
        }
        if (this.fotoProgressFile) {
          formData.append('foto_progress', this.fotoProgressFile);
        }
        if (this.fotoAfterFile) {
          formData.append('foto_after', this.fotoAfterFile);
        }

        await axios.put(`${apiEndpoint}${this.$route.params.id}`, formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        });

        var stat = '';

        if (apiEndpoint == '/api/program/draft/') {
          stat = 'draft';
        }else if (apiEndpoint == '/api/program/edit/') {
          stat = 'simpan';
        }else if (apiEndpoint == '/api/program/publish/') {
          stat = 'publish';
        }

        // Tampilkan sukses
        await Swal.fire({
          icon: 'success',
          title: 'Berhasil!',
          text: `Program berhasil di${stat}`,
          timer: 1500,
          showConfirmButton: false
        });

        this.$router.go(-1);
      } catch (error) {
        console.error('Error:', error);
        Swal.fire({
          icon: 'error',
          title: 'Oops...',
          text: 'Terjadi kesalahan saat menyimpan program',
          confirmButtonText: 'Tutup'
        });
      }
    },

    initMap() {
      // Hapus map yang lama jika ada
      if (this.map) {
        this.map.remove();
      }

      // Inisialisasi peta baru
      this.map = L.map('map').setView([-2.5489, 118.0149], 5);

      L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '© OpenStreetMap contributors'
      }).addTo(this.map);

      // Debug: cek nilai koordinat
      console.log('Koordinat:', {
        lat: this.user.latitude,
        lng: this.user.longitude
      });

      // Set marker jika koordinat ada
      if (this.user.latitude && this.user.longitude) {
        const lat = parseFloat(this.user.latitude);
        const lng = parseFloat(this.user.longitude);

        console.log('Parsed koordinat:', { lat, lng });

        if (!isNaN(lat) && !isNaN(lng) &&
            lat >= -90 && lat <= 90 &&
            lng >= -180 && lng <= 180) {

          // Tambahkan marker
          this.marker = L.marker([lat, lng]).addTo(this.map);

          // Zoom ke lokasi
          this.map.setView([lat, lng], 15);

          console.log('Marker ditambahkan pada:', { lat, lng });
        }
      }

      // Event click untuk update lokasi
      this.map.on('click', (e) => {
        const { lat, lng } = e.latlng;

        // Update koordinat
        this.user.latitude = lat.toFixed(6);
        this.user.longitude = lng.toFixed(6);

        // Update marker
        if (this.marker) {
          this.marker.setLatLng([lat, lng]);
        } else {
          this.marker = L.marker([lat, lng]).addTo(this.map);
        }

      });

      // Inisialisasi search provider
      this.searchProvider = new OpenStreetMapProvider();
    },

    // Tambahkan method untuk reset marker
    resetMarker() {
      if (this.marker) {
        this.marker.remove();
        this.marker = null;
      }
      this.user.latitude = '';
      this.user.longitude = '';
      this.map.setView([-2.5489, 118.0149], 5); // Reset view ke Indonesia
    },

    async handleSearch() {
      if (!this.searchQuery) return;

      try {
        const results = await this.searchProvider.search({ query: this.searchQuery });

        if (results && results.length > 0) {
          const { x: lng, y: lat } = results[0];

          // Update koordinat
          this.user.latitude = lat.toFixed(6);
          this.user.longitude = lng.toFixed(6);

          // Update marker dan view
          if (this.marker) {
            this.marker.setLatLng([lat, lng]);
          } else {
            this.marker = L.marker([lat, lng]).addTo(this.map);
          }

          this.map.setView([lat, lng], 15);

          // Reset search query
          this.searchQuery = '';
        } else {
          Swal.fire({
            icon: 'warning',
            title: 'Lokasi Tidak Ditemukan',
            text: 'Silakan coba dengan kata kunci lain',
            timer: 2000,
            showConfirmButton: false
          });
        }
      } catch (error) {
        console.error('Error searching location:', error);
        Swal.fire({
          icon: 'error',
          title: 'Error',
          text: 'Gagal mencari lokasi',
          timer: 2000,
          showConfirmButton: false
        });
      }
    },
    formatJumlahAnggaran(event) {
      // Hapus karakter non-digit
      this.user.jumlah_anggaran = event.target.value.replace(/\D/g, '');
    },
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
    },
    'user.latitude': {
      handler(newVal) {
        if (!newVal && this.marker) {
          this.resetMarker();
        }
      }
    },
    'user.longitude': {
      handler(newVal) {
        if (!newVal && this.marker) {
          this.resetMarker();
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

  beforeDestroy() {
    if (this.map) {
      this.map.remove();
      this.map = null;
    }
    if (this.marker) {
      this.marker = null;
    }
  }
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

/* Fix icon markers yang hilang */
.leaflet-default-icon-path {
  background-image: url(https://unpkg.com/leaflet@1.7.1/dist/images/marker-icon.png);
}

/* Optional: styling untuk kontainer peta */
#map {
  border: 1px solid #ddd;
  border-radius: 4px;
}

/* Tambahkan style untuk memastikan peta terlihat */
#map {
  min-height: 400px;
  width: 100%;
  z-index: 1;
}

/* Fix untuk marker yang tidak muncul */
.leaflet-marker-icon {
  width: 25px !important;
  height: 41px !important;
}

/* Fix untuk popup */
.leaflet-popup-content {
  margin: 13px;
  min-width: 150px;
}

.search-container {
  margin-bottom: 10px;
  z-index: 1000;
  width: 100%;
}
</style>

