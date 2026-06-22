<template>
  <div class="min-h-screen bg-slate-50 text-slate-900 font-sans antialiased selection:bg-emerald-100">

    <!-- ─── HERO ──────────────────────────────────────────────── -->
    <PageHero
      title="Product"
      subtitle="Our main product is TeL Pellita Bleached Kraft Pulp based on 100% planted Pellita trees."
      :breadcrumbs="[{ label: 'Home', to: '/' }, { label: 'Product' }]"
    />

    <!-- ─── PRODUCT DESCRIPTION ───────────────────────────────── -->
    <main id="product-target" class="container mx-auto px-6 lg:px-16 mt-4 relative z-20 pb-24 scroll-mt-10">

      <div v-if="loading" class="flex items-center justify-center py-24">
        <div class="w-6 h-6 border-2 border-emerald-600 border-t-transparent rounded-full animate-spin"></div>
      </div>

      <div v-else-if="pageDescription" class="max-w-5xl mx-auto">
        <div class="product-description" v-html="pageDescription">
        </div>
      </div>

      <div v-else class="text-center py-24 text-slate-400">
        <p class="text-5xl mb-4">📝</p>
        <p class="text-xl">Product description not available.</p>
      </div>
    </main>

    <!-- ─── SLIDER / FACILITY ─────────────────────────────────── -->
    <section class="container mx-auto px-6 lg:px-16 pb-24">
      <div class="relative rounded-[3rem] overflow-hidden shadow-2xl h-[520px] group bg-[#14532d]">
        <TransitionGroup name="fade">
          <div v-for="(slide, i) in slides" :key="i" v-show="currentSlide === i" class="absolute inset-0">
            <img :src="getImageUrl(slide.image_url)" class="w-full h-full object-contain opacity-80 p-4"
              @error="(e) => e.target.src='https://images.unsplash.com/photo-1586528116311-ad8dd3c8310d?q=80&w=1200'" />
            <div class="absolute inset-0 bg-gradient-to-r from-black/80 via-black/20 to-transparent"></div>
            <div class="absolute inset-y-0 left-0 w-full md:w-2/3 flex items-center p-12 md:p-24 z-20">
              <div class="max-w-xl">
                <Transition name="reveal-up">
                  <div v-if="currentSlide === i" class="space-y-5">
                    <span class="text-emerald-400 font-bold tracking-[0.3em] text-sm uppercase">Industrial Excellence</span>
                    <h2 class="text-4xl md:text-5xl font-serif font-bold text-white leading-tight">{{ slide.title }}</h2>
                    <p class="text-slate-300 text-lg font-light leading-relaxed">{{ slide.description }}</p>
                  </div>
                </Transition>
              </div>
            </div>
          </div>
        </TransitionGroup>

        <!-- Progress bar -->
        <div class="absolute bottom-0 left-0 h-1 bg-emerald-500 z-40 transition-none" :style="{ width: progress + '%' }"></div>

        <!-- Arrows -->
        <div class="absolute bottom-10 right-10 z-30 flex gap-3 opacity-0 group-hover:opacity-100 transition-opacity duration-500">
          <button @click="manualNav(prevSlide)" class="p-3.5 rounded-full bg-white/10 backdrop-blur-md text-white border border-white/20 hover:bg-emerald-600 transition-all cursor-pointer">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/></svg>
          </button>
          <button @click="manualNav(nextSlide)" class="p-3.5 rounded-full bg-white/10 backdrop-blur-md text-white border border-white/20 hover:bg-emerald-600 transition-all cursor-pointer">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
          </button>
        </div>
 
        <!-- Dots -->
        <div class="absolute bottom-10 left-12 md:left-24 z-30 flex gap-3">
          <button v-for="(_, i) in slides" :key="i" @click="manualNav(() => goToSlide(i))"
            class="h-1 transition-all duration-500 rounded-full cursor-pointer"
            :class="currentSlide === i ? 'w-12 bg-emerald-500' : 'w-4 bg-white/20 hover:bg-white/40'">
          </button>
        </div>
      </div>
    </section>

  </div>

  <section class="contact-section">
      <div class="contact-container">
        <div class="contact-header anim-item">
          <span class="contact-label">GET IN TOUCH</span>
          <h2 class="contact-title">Our Offices</h2>
        </div>

        <div class="contact-grid">
          <div class="contact-card anim-item">
            <div class="contact-image">
              <img src="/images/jakarta.jpeg" alt="Jakarta Office" />
            </div>
            <div class="contact-body">
              <h3 class="contact-city">Jakarta</h3>
              <p class="contact-address">Menara Astra 22nd floor – Zona D, Jalan Jenderal Sudirman Kav. 5-6 Kel. Karet Tengsin, Kec. Tanah Abang</p>
              <a class="contact-phone" href="tel:+622186656809">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72c.127.96.361 1.903.7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45c.907.339 1.85.573 2.81.7A2 2 0 0 1 22 16.92z"/></svg>
                +62 21 8665 6809 / 8665 6810
              </a>
            </div>
          </div>

          <div class="contact-card anim-item">
            <div class="contact-image">
              <img src="/images/lokasi pabrik.jpeg" alt="Mill Site" />
            </div>
            <div class="contact-body">
              <h3 class="contact-city">Mill Site</h3>
              <p class="contact-address">Desa Banuayu, Kec. Empat Petulai Dangku, Kab. Muara Enim, Sumatera Selatan</p>
              <a class="contact-phone" href="tel:+62713324150">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72c.127.96.361 1.903.7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45c.907.339 1.85.573 2.81.7A2 2 0 0 1 22 16.92z"/></svg>
                (62) (713) 324150 – 324160
              </a>
            </div>
          </div>

          <div class="contact-card anim-item">
            <div class="contact-image">
              <img src="/images/palembang.jpeg" alt="Palembang Office" />
            </div>
            <div class="contact-body">
              <h3 class="contact-city">Palembang</h3>
              <p class="contact-address">Ruko Blok I/29, Komplek PTC Mall. Jl. R. Soekamto Palembang 30114, Sumatera Selatan</p>
              <a class="contact-phone" href="tel:+62711382409">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72c.127.96.361 1.903.7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45c.907.339 1.85.573 2.81.7A2 2 0 0 1 22 16.92z"/></svg>
                (62) (711) 382409
              </a>
            </div>
          </div>

          <div class="contact-card anim-item">
            <div class="contact-image">
              <img src="/images/tarahan.jpeg" alt="Tarahan Port" />
            </div>
            <div class="contact-body">
              <h3 class="contact-city">Tarahan</h3>
              <p class="contact-address">Jl. Soekarno Hatta Km. 14, Batu Serampok Kel. Srengsem Kec. Panjang, Bandar Lampung</p>
              <a class="contact-phone" href="tel:+62721342311">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72c.127.96.361 1.903.7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45c.907.339 1.85.573 2.81.7A2 2 0 0 1 22 16.92z"/></svg>
                (62) (721) 34231, 31318
              </a>
            </div>
          </div>
        </div>
      </div>
    </section>

    <section class="footer-info-section">
      <div class="footer-info-container">
        <div class="footer-info-grid">

          <div class="footer-info-col anim-item">
            <p class="footer-about-text">
              PT. Tanjungenim Lestari Pulp and Paper (PT. TEL), is one of the most exciting pulp mills in Indonesia today and the only pulp mill in the world to produce high-quality, bleached-hardwood Kraft pulp with 100 percent plantation grown Acacia mangium and Eucalyptus Pellita trees.
            </p>
          </div>

          <div class="footer-info-col footer-logos-col anim-item">
            <img src="/images/logosatu.jpeg" alt="Yayasan Pendidikan Tanjungenim Lestari - SMP, SD, PGTK Lematang Lestari" class="footer-logo-combined" />
          </div>

          <div class="footer-info-col footer-contact-col anim-item">
            <div class="footer-contact-row">
              <svg class="footer-contact-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 1 1 18 0z"/><circle cx="12" cy="10" r="3"/></svg>
              <span>Kab. Muara Enim, Sumatera Selatan.</span>
            </div>
            <div class="footer-contact-row">
              <svg class="footer-contact-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72c.127.96.361 1.903.7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45c.907.339 1.85.573 2.81.7A2 2 0 0 1 22 16.92z"/></svg>
              <span>(+62) 713-324-150</span>
            </div>
            <div class="footer-contact-row">
              <svg class="footer-contact-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 6v6l4 2"/></svg>
              <span>Mon-Fri: 8:00 – 17:00</span>
            </div>
          </div>

        </div>
      </div>
    </section>
    
  <footer class="site-footer">
      <div class="footer-container">
        <div class="footer-left-content"></div>
        <div class="footer-copyright">
          <p>Copyright 2026 PT TELPP. All right reserved.</p>
        </div>
      </div>
    </footer>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import api from '../services/api'
import PageHero from '../components/PageHero.vue'

const BASE_URL = import.meta.env.VITE_BASE_URL || 'http://localhost:8080'

const pageDescription = ref('')
const loading  = ref(true)

onMounted(async () => {
  try {
    const [pageRes, slideRes] = await Promise.all([
      api.get('/product-page'),
      api.get('/product-slides')
    ])
    pageDescription.value = pageRes.data?.data?.description || ''
    if (slideRes.data?.data) {slides.value = slideRes.data.data}
  } catch {
    pageDescription.value = ''
  } finally {
    loading.value = false
  }
  startTimer()
})
onUnmounted(() => stopTimer())

// ─── Slider ───────────────────────────────────────────────────
const slides = ref([])
const currentSlide = ref(0)
const progress = ref(0)
let timer = null, progressTimer = null

const getImageUrl = (path) => {
  if (!path) {return ''}
  if (path.startsWith('http')) {return path}
  return `${BASE_URL}/${path.replace(/^\//, '')}`
}

const startTimer = () => {
  timer = setInterval(nextSlide, 6000)
  const step = 100 / (6000 / 10)
  progressTimer = setInterval(() => { progress.value = Math.min(progress.value + step, 100) }, 10)
}
const stopTimer  = () => { clearInterval(timer); clearInterval(progressTimer) }
const nextSlide  = () => { if (!slides.value.length) {return}; currentSlide.value = (currentSlide.value + 1) % slides.value.length; progress.value = 0 }
const prevSlide  = () => { if (!slides.value.length) {return}; currentSlide.value = (currentSlide.value - 1 + slides.value.length) % slides.value.length; progress.value = 0 }
const goToSlide  = (i) => { currentSlide.value = i; progress.value = 0 }
const manualNav  = (fn) => { stopTimer(); fn(); startTimer() }
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@300;400;600;700&family=Playfair+Display:wght@700&display=swap');
.font-serif { font-family: 'Playfair Display', serif; }

.list-fade-enter-active { transition: all 0.6s cubic-bezier(0.34,1.56,0.64,1); }
.list-fade-enter-from   { opacity: 0; transform: scale(0.85) translateY(30px); }
.reveal-up-enter-active   { transition: all 1s cubic-bezier(0.22,1,0.36,1); }
.reveal-up-enter-from     { opacity: 0; transform: translateY(40px); }
.reveal-down-enter-active { transition: all 1s cubic-bezier(0.22,1,0.36,1); }
.reveal-down-enter-from   { opacity: 0; transform: translateY(-40px); }
.fade-enter-active, .fade-leave-active { transition: opacity 1.2s ease-in-out; }
.fade-enter-from, .fade-leave-to       { opacity: 0; }

@keyframes subtle-zoom { 0% { transform:scale(1); } 100% { transform:scale(1.08); } }
.animate-subtle-zoom { animation: subtle-zoom 20s infinite alternate ease-in-out; }
@keyframes bounce { 0%,100%{transform:translateY(-5%);} 50%{transform:none;} }
.animate-bounce-slow { animation: bounce 2s infinite; }

.contact-container {
  max-width: 1140px;
  margin: 0 auto;
  padding: 0 40px;
}

.contact-header {
  text-align: center;
  margin-bottom: 48px;
}

.contact-label {
  color: #5F9E42;
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 0.25em;
  text-transform: uppercase;
  display: block;
  margin-bottom: 10px;
}

.contact-title {
  font-size: 32px;
  font-weight: 800;
  color: #18301f;
  letter-spacing: -0.02em;
}

.contact-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 24px;
}

@media (max-width: 1024px) {
  .contact-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 640px) {
  .contact-grid {
    grid-template-columns: 1fr;
  }
}

.contact-card {
  background: #ffffff;
  border: 1px solid #dcecd2;
  border-radius: 20px;
  padding: 28px 24px;
  transition: all 0.4s cubic-bezier(0.16, 1, 0.3, 1);
}

.contact-card:hover {
  transform: translateY(-6px);
  border-color: #a9d68c;
  box-shadow: 0 16px 32px rgba(95, 158, 66, 0.12);
}

.contact-icon {
  width: 44px;
  height: 44px;
  border-radius: 12px;
  background: #eaf6e1;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 18px;
  color: #5F9E42;
}

.contact-icon svg {
  width: 22px;
  height: 22px;
}

.contact-city {
  font-size: 17px;
  font-weight: 700;
  color: #18301f;
  margin-bottom: 10px;
}

.contact-address {
  font-size: 13px;
  color: #6b7b66;
  line-height: 1.6;
  margin-bottom: 16px;
  font-weight: 400;
}

.contact-phone {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13.5px;
  font-weight: 600;
  color: #3b6d11;
  text-decoration: none;
  transition: color 0.3s ease;
}

.contact-phone:hover {
  color: #27500a;
}

.contact-phone svg {
  width: 15px;
  height: 15px;
  flex-shrink: 0;
}

.contact-section {
  background: linear-gradient(180deg, #ffffff 0%, #eef7e9 100%);
  padding: 80px 0 70px;
}

.contact-container {
  max-width: 1140px;
  margin: 0 auto;
  padding: 0 40px;
}

.contact-header {
  text-align: center;
  margin-bottom: 48px;
}

.contact-label {
  color: #5F9E42;
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 0.25em;
  text-transform: uppercase;
  display: block;
  margin-bottom: 10px;
}

.contact-title {
  font-size: 32px;
  font-weight: 800;
  color: #18301f;
  letter-spacing: -0.02em;
}

.contact-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 24px;
}

@media (max-width: 1024px) {
  .contact-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 640px) {
  .contact-grid {
    grid-template-columns: 1fr;
  }
}

.contact-card {
  background: #ffffff;
  border: 1px solid #dcecd2;
  border-radius: 20px;
  overflow: hidden;
  transition: all 0.4s cubic-bezier(0.16, 1, 0.3, 1);
}

.contact-card:hover {
  transform: translateY(-6px);
  border-color: #a9d68c;
  box-shadow: 0 16px 32px rgba(95, 158, 66, 0.15);
}

.contact-image {
  width: 100%;
  aspect-ratio: 4 / 3;
  overflow: hidden;
  background: #eaf6e1;
}

.contact-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
  transition: transform 0.6s cubic-bezier(0.16, 1, 0.3, 1);
}

.contact-card:hover .contact-image img {
  transform: scale(1.06);
}

.contact-body {
  padding: 20px 18px 22px;
}

.contact-city {
  font-size: 18px;
  font-weight: 800;
  color: #1a4d8f;
  margin-bottom: 12px;
  letter-spacing: -0.01em;
}

.contact-address {
  font-size: 13px;
  color: #5a6b56;
  line-height: 1.6;
  margin-bottom: 16px;
  font-weight: 400;
}

.contact-phone {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13.5px;
  font-weight: 500;
  color: #2c3a2a;
  text-decoration: none;
  transition: color 0.3s ease;
}

.contact-phone:hover {
  color: #3b6d11;
}

.contact-phone svg {
  width: 15px;
  height: 15px;
  flex-shrink: 0;
  color: #5F9E42;
}
/* Responsif untuk Layar Handphone (Mobile) */
@media (max-width: 768px) {
  .footer-container {
    flex-direction: column;
    gap: 15px;
    padding: 0 20px;
  }
  
  .footer-copyright {
    text-align: center; /* Di HP otomatis pindah ke tengah bawah agar rapi */
    width: 100%;
  }
  
  .footer-copyright p {
    font-size: 12px; /* Ukuran sedikit disesuaikan untuk layar kecil */
  }
}

.footer-info-section {
  background: #ffffff;
  border-top: 1px solid #e4efdb;
  padding: 56px 0;
}

.footer-info-container {
  max-width: 1140px;
  margin: 0 auto;
  padding: 0 40px;
}
.footer-info-grid {
  display: grid;
  grid-template-columns: 1fr 1.1fr 0.9fr;
  gap: 48px;
  align-items: start;
}

.footer-info-col {
  display: flex;
  flex-direction: column;
}

.footer-about-text {
  margin: 0;
  font-size: 14px;
  color: #7a8a76;
  line-height: 1.75;
  font-weight: 400;
}

.footer-logos-col {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0;
  margin: 0;
  padding: 0;
}

.footer-logos-col {
  display: flex;
  align-items: flex-start;
  justify-content: center;
  margin: 0;
  padding: 0;
}

.footer-logo-combined {
  width: 100%;
  max-width: 380px;
  height: auto;
  object-fit: contain;
  display: block;
}

.footer-contact-col {
  display: flex;
  flex-direction: column;
  gap: 20px;
  max-width: 280px;
  margin: 0;
  padding: 0;
}

.footer-contact-row {
  display: flex;
  align-items: center;
  gap: 14px;
  font-size: 14px;
  color: #5a6b56;
  font-weight: 400;
  line-height: 1.4;
  margin: 0;
}
.footer-contact-icon {
  width: 19px;
  height: 19px;
  flex-shrink: 0;
  color: #5F9E42;
}

/* ============ FOOTER ============ */
.site-footer {
  width: 100%;
  /* Warna hijau tua disesuaikan dengan identitas PT TEL pada gambar */
  background-color: #5F9E42; 
  padding: 25px 0; /* Memberikan ruang tinggi baris yang pas */
  margin-top: 40px; /* Jarak pemisah dari konten di atasnya */
}

.footer-container {
  display: flex;
  align-items: center;
  justify-content: center; /* DIUBAH: Membuat konten berkumpul di tengah */
  max-width: 1140px; /* Menyelaraskan lebar dengan konten atas website */
  margin: 0 auto;
  padding: 0 40px; /* Jarak aman agar teks tidak terlalu menempel ke pinggir layar */
}

/* Kolom Kiri Footer */
.footer-left-content {
  display: none; /* DIUBAH: Disembunyikan karena kosong, agar tidak memakan space */
}

/* Kolom Kanan Footer (Teks Copyright) */
.footer-copyright {
  text-align: center; /* DIUBAH: Memastikan teks rata tengah */
  width: 100%; /* Memastikan block mengambil ruang penuh untuk fungsi centering */
}

.footer-copyright p {
  font-family: Arial, sans-serif;
  font-size: 13px; /* Ukuran teks kecil, tipis, dan normal untuk footer */
  color: #ffffff;  /* Warna teks putih bersih agar kontras dengan background hijau */
  margin: 0;
  font-weight: normal;
  letter-spacing: 0.3px;
}
@media (max-width: 768px) {
  .footer-info-grid {
    grid-template-columns: 1fr !important;
    gap: 32px !important;
  }
  .footer-contact-col {
    max-width: 100% !important;
  }
  .footer-info-container,
  .contact-container {
    padding: 0 20px !important;
  }
  .contact-section {
    padding: 48px 0 40px;
  }
  .footer-info-section {
    padding: 36px 0;
  }
}
</style>

<style>
.product-description { color: #374151; font-size: 16px; line-height: 1.75; padding: 0; }
.product-description p { margin-bottom: 1em; }
.product-description h1 { font-size: 1.5em; font-weight: 700; margin: 0.8em 0 0.4em; color: #111827; }
.product-description h2 { font-size: 1.25em; font-weight: 600; margin: 0.8em 0 0.4em; color: #111827; }
.product-description h3 { font-size: 1.1em; font-weight: 600; margin: 0.8em 0 0.4em; color: #111827; }
.product-description ul { list-style-type: disc; padding-left: 1.5em; margin: 0.5em 0; }
.product-description ol { list-style-type: decimal; padding-left: 1.5em; margin: 0.5em 0; }
.product-description li { margin: 0.25em 0; }
.product-description strong { font-weight: 700; color: #111827; }
.product-description em { font-style: italic; }
.product-description pre { background: #1e293b; color: #e2e8f0; padding: 12px 16px; border-radius: 8px; font-family: monospace; font-size: 0.85em; overflow-x: auto; margin: 0.75em 0; }
.product-description code { background: #f1f5f9; padding: 2px 6px; border-radius: 4px; font-family: monospace; font-size: 0.9em; }
.product-description pre code { background: transparent; padding: 0; }
.product-description img { max-width: 100%; border-radius: 8px; margin: 1em 0; }
.product-description blockquote { border-left: 3px solid #10b981; padding-left: 12px; margin: 0.75em 0; color: #6b7280; }
</style>