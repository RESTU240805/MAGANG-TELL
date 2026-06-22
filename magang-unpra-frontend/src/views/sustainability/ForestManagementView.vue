<template>
  <div class="bg-[#F8F9FA] text-zinc-900 font-sans antialiased overflow-x-hidden selection:bg-green-500 selection:text-white">

    <PageHero
      title="Forest Management"
      subtitle="Sustainable and responsible forest management for a greener future in South Sumatra."
      :breadcrumbs="[{ label: 'Home', to: '/' }, { label: 'Sustainability' }, { label: 'Forest Management' }]"
    />

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
        <div class="relative rounded-3xl overflow-hidden shadow-2xl bg-zinc-100 aspect-[16/9] anim-item group">
          <img v-for="(slide, i) in heroSlides" :key="i"
            :src="slide" alt="Forest Management"
            class="absolute inset-0 w-full h-full object-cover"
            :style="{ opacity: currentSlide === i ? 1 : 0, transition: 'opacity 1s ease-in-out' }" />
          <div class="absolute inset-0 bg-gradient-to-t from-black/20 to-transparent"></div>
          <button @click="manualHeroNav(prevHeroSlide)"
            class="absolute left-4 top-1/2 -translate-y-1/2 w-10 h-10 rounded-full bg-white/20 backdrop-blur-md text-white border border-white/30 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity cursor-pointer hover:bg-white/40">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/></svg>
          </button>
          <button @click="manualHeroNav(nextHeroSlide)"
            class="absolute right-4 top-1/2 -translate-y-1/2 w-10 h-10 rounded-full bg-white/20 backdrop-blur-md text-white border border-white/30 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity cursor-pointer hover:bg-white/40">
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
  </div>
</template>

<script setup>
import PageHero from '../../components/PageHero.vue'
import { ref, onMounted, onUnmounted, nextTick } from 'vue'

let observer = null
let heroTimer = null

const currentSlide = ref(0)
const heroSlides = [
  '/images/hutan hd.jpeg',
  '/images/hutanhutan.jpeg',
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