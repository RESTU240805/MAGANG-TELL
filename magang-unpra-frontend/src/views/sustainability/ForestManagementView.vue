<template>
  <div class="bg-[#F8F9FA] text-zinc-900 font-sans antialiased overflow-x-hidden selection:bg-green-500 selection:text-white">
    
    <section class="relative min-h-[90vh] md:h-[95vh] flex flex-col justify-between bg-gradient-to-br from-[#041a0f] to-zinc-950 overflow-hidden pb-12">
      
      <div 
        class="absolute inset-0 bg-cover bg-center scale-110"
        style="background-image: url('/images/hutan hd.jpeg')"
      ></div>
      
      <div class="absolute inset-0 bg-gradient-to-t from-black/80 via-black/30 to-black/10"></div>

      <div></div>

      <div class="relative z-10 max-w-6xl mx-auto px-6 md:px-12 w-full mt-24 lg:mt-32">
        <span class="text-green-500 text-xs font-bold tracking-[0.25em] uppercase block mb-4 opacity-0 anim-hero-reveal" style="animation-delay: 300ms">
          SUSTAINABILITY 
        </span>
        <h1 class="text-5xl md:text-7xl font-extrabold text-white tracking-tight leading-[1.05] mb-6 opacity-0 anim-hero-reveal" style="animation-delay: 550ms">
          Forest<br />
          <span class="text-green-500 font-black">Management</span>
        </h1>
        <p class="text-zinc-300 text-sm md:text-base max-w-md leading-relaxed mb-8 opacity-0 anim-hero-reveal" style="animation-delay: 800ms">
          Sustainable and responsible forest management for a greener future in South Sumatra.
        </p>
      </div>

      <div class="relative z-10 max-w-6xl mx-auto px-6 md:px-12 w-full mt-20 flex flex-col sm:flex-row items-start sm:items-center justify-between gap-6 opacity-0 anim-hero-reveal" style="animation-delay: 1050ms">
        <div class="flex items-center gap-3 group text-zinc-400 cursor-pointer">
          <div class="w-5 h-8 border-2 border-zinc-600 rounded-full flex justify-center p-1 group-hover:border-green-500 transition-colors duration-500">
            <div class="w-1 h-1.5 bg-green-400 rounded-full animate-[scrollDrop_2s_ease-in-out_infinite]"></div>
          </div>
          <span class="text-[10px] font-bold tracking-[0.2em] uppercase text-zinc-500 group-hover:text-zinc-200 transition-colors duration-500">Scroll To Explore</span>
        </div>
      </div>
    </section>

    <section class="py-24 bg-gradient-to-b from-[#F8F9FA] to-white relative">
      <div class="max-w-6xl mx-auto px-6 md:px-12">
        
        <div class="text-center mb-20 anim-reveal-up">
          <span class="text-green-600 text-xs font-bold tracking-widest uppercase block mb-2">RAW MATERIALS</span>
          <h2 class="text-3xl md:text-4xl font-extrabold text-zinc-900 tracking-tight">Types of Wood Used</h2>
          <p class="text-zinc-500 text-sm mt-3 max-w-xl mx-auto leading-relaxed">
            Pulp is produced from two types of wood sourced from sustainable industrial plantations.
          </p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
          
          <div class="bg-[#0b2416] rounded-[28px] overflow-hidden group shadow-xl hover:shadow-[0_20px_40px_rgba(11,36,22,0.15)] hover:-translate-y-2 transition-all duration-700 ease-[cubic-bezier(0.16,1,0.3,1)] flex flex-col justify-between anim-reveal-left">
            <div class="p-8 pb-8">
              <div class="relative rounded-2xl overflow-hidden aspect-[16/10] mb-6 bg-zinc-950">
                <img src="/images/accsia.jpg" alt="Acacia Mangium" class="w-full h-full object-cover scale-100 group-hover:scale-105 transition-transform duration-[1200ms] ease-[cubic-bezier(0.16,1,0.3,1)] opacity-90 group-hover:opacity-100" />
                <div class="absolute inset-0 bg-gradient-to-t from-black/40 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-700"></div>
              </div>
              <span class="inline-block bg-green-500/10 text-green-400 text-[10px] font-bold tracking-wider uppercase px-3 py-1 rounded-md mb-4 border border-green-500/20">
                Initial Product
              </span>
              <h3 class="text-xl font-bold text-white mb-3 group-hover:text-green-400 transition-colors duration-300">Acacia Mangium (AM)</h3>
              <p class="text-zinc-400 text-xs md:text-sm leading-relaxed font-light">
                The main product since the start of TEL's plant operation in 1999. Acacia Mangium Pulp served as the foundation of production.
              </p>
            </div>
          </div>

          <div class="bg-[#0b2416] rounded-[28px] overflow-hidden group shadow-xl hover:shadow-[0_20px_40px_rgba(11,36,22,0.15)] hover:-translate-y-2 transition-all duration-700 ease-[cubic-bezier(0.16,1,0.3,1)] flex flex-col justify-between anim-reveal-right">
            <div class="p-8 pb-8">
              <div class="relative rounded-2xl overflow-hidden aspect-[16/10] mb-6 bg-zinc-950">
                <img src="/images/eccalyptus.jpg" alt="Eucalyptus Pellita" class="w-full h-full object-cover scale-100 group-hover:scale-105 transition-transform duration-[1200ms] ease-[cubic-bezier(0.16,1,0.3,1)] opacity-90 group-hover:opacity-100" />
                <div class="absolute inset-0 bg-gradient-to-t from-black/40 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-700"></div>
              </div>
              <span class="inline-block bg-green-500/10 text-green-400 text-[10px] font-bold tracking-wider uppercase px-3 py-1 rounded-md mb-4 border border-green-500/20">
                Main Product Now
              </span>
              <h3 class="text-xl font-bold text-white mb-3 group-hover:text-green-400 transition-colors duration-300">Eucalyptus Pellita (EP)</h3>
              <p class="text-zinc-400 text-xs md:text-sm leading-relaxed font-light">
                Through continuous innovation and development, Eucalyptus Pellita Pulp has now become the primary raw material used by TEL.
              </p>
            </div>
          </div>

        </div>
      </div>
    </section>

    <section class="py-24 bg-white border-t border-zinc-100 overflow-hidden">
      <div class="max-w-6xl mx-auto px-6 md:px-12">

        <!-- Text (Top) -->
        <div class="max-w-3xl mx-auto text-center mb-16 anim-item">
          <span class="text-green-600 text-xs font-bold tracking-widest uppercase block mb-3">OUR APPROACH</span>
          <h2 class="text-3xl md:text-4xl font-extrabold text-zinc-900 tracking-tight leading-tight mb-6">
            High Quality Pulp from <br />
            <span class="text-green-600 font-black">100% Sustainable Wood</span>
          </h2>
          <div class="space-y-5 text-zinc-500 text-sm leading-relaxed font-normal text-left">
            <p>PT Tanjungenim Lestari Pulp and Paper (TEL) produces high quality pulp from 100% wood based on industrial plantations sourced from a group company, PT Musi Hutan Persada (MHP), a highly dedicated Industrial Plantation Company with 296,400 ha of concession area in South Sumatra. MHP runs its operations with Sustainable and Sustainable Forest Management.</p>
            <p>Pulp is produced from two types of wood, namely Acacia Mangium (AM) and Eucalyptus Pellita (EP). Since the start of the plant's operation in 1999, the main product is the Acacia Mangium Pulp and then through continuous efforts, TEL is working with MHP to develop new alternative raw materials, until now Eucalyptus Pellita Pulp has become the main product.</p>
            <p>Timber delivery from the Industrial Plantation area to the Factory Site continues 365 days per year by truck through a special road network built by MHP. TEL sets standards for the size and quality of logs that must be met by MHP.</p>
          </div>
        </div>

        <!-- Image Slider (Bottom) -->
        <div class="relative rounded-3xl overflow-hidden shadow-2xl bg-zinc-100 aspect-[16/9] anim-item">
          <img v-for="(slide, i) in heroSlides" :key="i"
            :src="slide" alt="Forest Management"
            class="absolute inset-0 w-full h-full object-cover"
            :style="{ opacity: currentSlide === i ? 1 : 0, transition: 'opacity 1s ease-in-out' }" />
          <div class="absolute inset-0 bg-gradient-to-t from-black/20 to-transparent"></div>
          <button @click="manualHeroNav(prevHeroSlide)"
            class="absolute left-4 top-1/2 -translate-y-1/2 w-10 h-10 rounded-full bg-white/20 backdrop-blur-md text-white border border-white/30 flex items-center justify-center opacity-0 hover:opacity-100 transition-opacity cursor-pointer hover:bg-white/40">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/></svg>
          </button>
          <button @click="manualHeroNav(nextHeroSlide)"
            class="absolute right-4 top-1/2 -translate-y-1/2 w-10 h-10 rounded-full bg-white/20 backdrop-blur-md text-white border border-white/30 flex items-center justify-center opacity-0 hover:opacity-100 transition-opacity cursor-pointer hover:bg-white/40">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
          </button>
          <div class="absolute bottom-4 left-1/2 -translate-x-1/2 flex gap-2 z-10">
            <button v-for="(_, i) in heroSlides" :key="i" @click="manualHeroNav(() => goToHeroSlide(i))"
              class="h-1.5 rounded-full transition-all duration-500 cursor-pointer"
              :class="currentSlide === i ? 'w-8 bg-white' : 'w-3 bg-white/40 hover:bg-white/60'"></button>
          </div>
        </div>

      </div>
    </section>

    <section class="py-12 bg-white">
      <div class="max-w-6xl mx-auto px-6 md:px-12">
        <div class="relative bg-gradient-to-r from-[#031d10] to-[#0a351d] rounded-3xl overflow-hidden p-8 md:p-12 shadow-xl flex flex-col md:flex-row md:items-center justify-between gap-8 anim-reveal-up">
          <div class="absolute right-0 bottom-0 top-0 w-1/2 opacity-20 md:opacity-30 pointer-events-none select-none">
            <img src="/images/lokasi pabrik.jpeg" class="w-full h-full object-cover mix-blend-overlay scale-105 transition-transform duration-1000" />
          </div>
          <div class="relative z-10 max-w-xl">
            <span class="text-green-400 text-[10px] font-bold tracking-widest uppercase block mb-2">TOGETHER FOR A GREEN FUTURE</span>
            <h2 class="text-2xl md:text-3xl font-bold text-white tracking-tight leading-tight">
              Building a Sustainable Future <br />Through Responsible Forest Management
            </h2>
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

  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, nextTick } from 'vue'

let observer = null
let heroTimer = null

const currentSlide = ref(0)
const heroSlides = [
  '/images/hutan hd.jpeg',
  '/images/hutan3.jpeg',
  '/images/mobil kayu.jpeg',
  '/images/pohon.jpeg'
]

const nextHeroSlide = () => { currentSlide.value = (currentSlide.value + 1) % heroSlides.length }
const prevHeroSlide = () => { currentSlide.value = (currentSlide.value - 1 + heroSlides.length) % heroSlides.length }
const goToHeroSlide = (i) => { currentSlide.value = i }
const startHeroTimer = () => { heroTimer = setInterval(nextHeroSlide, 5000) }
const stopHeroTimer = () => { clearInterval(heroTimer) }
const manualHeroNav = (fn) => { stopHeroTimer(); fn(); startHeroTimer() }

onMounted(async () => {
  await nextTick()
  initObserver()
  startHeroTimer()
})

onUnmounted(() => {
  if (observer) {observer.disconnect()}
  stopHeroTimer()
})

const initObserver = () => {
  observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.add('anim-active')
      } else {
        entry.target.classList.remove('anim-active')
      }
    })
  }, { 
    threshold: 0.08,
    rootMargin: '0px 0px -40px 0px'
  })

  document.querySelectorAll('.anim-reveal-up, .anim-reveal-left, .anim-reveal-right, .anim-item').forEach(el => {
    observer.observe(el)
  })
}
</script>

<style scoped>
@keyframes superZoomOut {
  0% { transform: scale(1.15); opacity: 0; filter: blur(4px); }
  100% { transform: scale(1); opacity: 1; filter: blur(0); }
}

@keyframes elegantReveal {
  0% { opacity: 0; transform: translateY(40px); }
  100% { opacity: 1; transform: translateY(0); }
}

@keyframes scrollDrop {
  0% { opacity: 0; transform: translateY(-8px); }
  30% { opacity: 1; }
  60% { opacity: 0; transform: translateY(10px); }
  100% { opacity: 0; }
}

.anim-hero-reveal {
  animation: elegantReveal 1.6s cubic-bezier(0.16, 1, 0.3, 1) forwards;
}

.anim-reveal-up {
  opacity: 0;
  transform: translateY(60px);
  transition: opacity 1.4s cubic-bezier(0.16, 1, 0.3, 1), transform 1.4s cubic-bezier(0.16, 1, 0.3, 1);
  will-change: transform, opacity;
}

.anim-reveal-left {
  opacity: 0;
  transform: translateX(-50px) translateY(10px);
  transition: opacity 1.5s cubic-bezier(0.16, 1, 0.3, 1), transform 1.5s cubic-bezier(0.16, 1, 0.3, 1);
  will-change: transform, opacity;
}

.anim-reveal-right {
  opacity: 0;
  transform: translateX(50px) translateY(10px);
  transition: opacity 1.5s cubic-bezier(0.16, 1, 0.3, 1), transform 1.5s cubic-bezier(0.16, 1, 0.3, 1);
  will-change: transform, opacity;
}

.anim-active, .anim-item.anim-active {
  opacity: 1 !important;
  transform: translate(0, 0) !important;
}

.anim-item {
  opacity: 0;
  transform: translateY(30px);
  transition: opacity 0.7s cubic-bezier(0.22, 1, 0.36, 1), transform 0.7s cubic-bezier(0.22, 1, 0.36, 1);
}

@keyframes smoothFloat {
  0%, 100% { transform: translateY(0) rotate(0deg); }
  50% { transform: translateY(-10px) rotate(0.5deg); }
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
</style>