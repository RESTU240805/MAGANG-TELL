<template>

  <!-- ── HERO ── -->
  <section class="relative h-[400px] flex items-end bg-cover bg-center overflow-hidden"
    style="background-image: url('/safety/1.png')">
    <div class="absolute inset-0 bg-gradient-to-t from-black/85 via-black/45 to-transparent"></div>

    <div class="absolute top-10 left-10 flex items-center gap-2 opacity-0 anim-fade-right" style="animation-delay: 0.1s">
      <span class="w-1 h-5 bg-green-400 rounded-full"></span>
      <span class="text-green-400 text-xs font-bold tracking-[3px] uppercase">Sustainability</span>
    </div>

    <div class="relative z-10 max-w-5xl mx-auto px-10 pb-14 w-full">
      <h1 class="text-5xl font-black text-white leading-tight mb-4 opacity-0 anim-fade-up" style="animation-delay: 0.2s">
        Safety & Health <span class="text-green-400">Office</span>
      </h1>
      <p class="text-gray-300 text-sm max-w-xl leading-relaxed opacity-0 anim-fade-up" style="animation-delay: 0.3s">
        Committed to implementing Occupational Safety and Health across all factory areas for every employee, contractor and supplier.
      </p>
      <div class="flex items-center gap-2 mt-5 text-xs text-gray-400 opacity-0 anim-fade-up" style="animation-delay: 0.4s">
        <RouterLink to="/" class="hover:text-white cursor-pointer transition">Home</RouterLink>
        <span class="text-gray-600">/</span>
        <span class="hover:text-white cursor-pointer transition">Sustainability</span>
        <span class="text-gray-600">/</span>
        <span class="text-green-400 font-semibold">Safety & Health Office</span>
      </div>
    </div>
  </section>

  <!-- ── K3 POLICY ── -->
  <section class="py-20 bg-white">
    <div class="max-w-5xl mx-auto px-10">

      <div class="flex items-center gap-4 mb-8 anim-item">
        <div class="w-12 h-12 rounded-2xl bg-[#14532d] flex items-center justify-center text-2xl flex-shrink-0">
          🛡️
        </div>
        <div>
          <span class="inline-block bg-green-100 text-green-700 text-xs font-bold tracking-widest uppercase px-3 py-1 rounded-full mb-1">
            K3 Policy
          </span>
          <h2 class="text-3xl font-black text-gray-900">Kebijakan K3</h2>
        </div>
      </div>

      <div class="grid grid-cols-5 gap-10">

        <div class="col-span-3 space-y-4 text-gray-500 text-sm leading-relaxed anim-slide-right">
          <p>Tanjungenim Lestari Pulp and Paper is fully committed to implementing Occupational Safety and Health in all factory areas by each of our employees, contractors and suppliers.</p>
          <p>The company believes that all employees are valuable assets for the company. This policy is a commitment in the field of Occupational Safety and Health to support all company activities with high productivity without creating risks to Occupational Safety and Health for employees and related parties.</p>
          <p>To ensure the efficient implementation of the management system, we have implemented an Occupational Safety and Health management system – <strong class="text-gray-700">SMK3 and OHSAS 18001</strong> with a spirit of continuous improvement and compliance with government regulations, the Company's Occupational Safety and Health Policy and other safety requirements.</p>
          <p>The company's attention is realized by providing training, programs and all the resources needed to ensure compliance with Occupational Safety and Health requirements.</p>
        </div>

        <div class="col-span-2 space-y-3 anim-slide-left">
          <div v-for="card in policyCards" :key="card.title"
            class="bg-green-50 border border-green-100 rounded-2xl p-5 hover:border-green-300 hover:bg-green-100/50 transition">
            <div class="text-xl mb-2">{{ card.icon }}</div>
            <p class="text-gray-900 font-bold text-sm mb-1">{{ card.title }}</p>
            <p class="text-gray-400 text-xs leading-relaxed">{{ card.desc }}</p>
          </div>
        </div>
      </div>

      <div class="mt-8 border-l-4 border-[#14532d] pl-5 bg-green-50 py-4 pr-5 rounded-r-xl anim-item">
        <p class="text-green-800 text-sm font-medium leading-relaxed">
          ⚠️ Occupational Safety and Health Policy of PT. Tanjungenim Lestari Pulp and Paper must be obeyed and followed by every employee, including contractors and suppliers seriously.
        </p>
      </div>
    </div>
  </section>

  <!-- ── KOMITE K3 ── -->
  <section class="py-20 bg-[#14532d]">
    <div class="max-w-5xl mx-auto px-10">

      <div class="flex items-center gap-4 mb-10 anim-item">
        <div class="w-12 h-12 rounded-2xl bg-green-500 flex items-center justify-center text-2xl flex-shrink-0">
          👥
        </div>
        <div>
          <span class="inline-block bg-green-800 text-green-200 text-xs font-bold tracking-widest uppercase px-3 py-1 rounded-full mb-1">
            Komite K3
          </span>
          <h2 class="text-3xl font-black text-white">P2K3 Committee</h2>
        </div>
      </div>

      <div class="grid grid-cols-2 gap-6">
        <div class="bg-white/10 backdrop-blur-sm rounded-2xl p-6 border border-white/10 hover:bg-white/15 transition anim-item">
          <div class="w-9 h-9 rounded-xl bg-green-500/30 flex items-center justify-center text-base mb-4">🏢</div>
          <h3 class="font-bold text-white text-sm mb-2">Safety Office & P2K3</h3>
          <p class="text-green-200/60 text-sm leading-relaxed">
            PT. Tanjungenim Lestari Pulp and Paper has a special section that accommodates the implementation of the K3 program, namely the Safety Office and the K3 Committee or the Occupational Safety and Health Management Committee (P2K3) consisting of management lines and employee representatives in each Department / Section.
          </p>
        </div>
        <div class="bg-white/10 backdrop-blur-sm rounded-2xl p-6 border border-white/10 hover:bg-white/15 transition anim-item" style="animation-delay: 0.1s">
          <div class="w-9 h-9 rounded-xl bg-green-500/30 flex items-center justify-center text-base mb-4">📋</div>
          <h3 class="font-bold text-white text-sm mb-2">Legal Authority</h3>
          <p class="text-green-200/60 text-sm leading-relaxed">
            The tasks and authority of P2K3 have been decided in the Decree of the Head of the South Sumatra Province Manpower and Transmigration Office No.254 / SK / Nakertrans / 2016. P2K3 in accordance with regulations of the Minister of Manpower has the duty to provide an output of consideration whether requested or not requested by management regarding aspects of occupational safety and health.
          </p>
        </div>
      </div>
    </div>
  </section>

  <!-- ── K3 TARGET ── -->
  <section class="py-20 bg-white">
    <div class="max-w-5xl mx-auto px-10">

      <div class="flex items-center gap-4 mb-4 anim-item">
        <div class="w-12 h-12 rounded-2xl bg-[#14532d] flex items-center justify-center text-2xl flex-shrink-0">
          🎯
        </div>
        <div>
          <span class="inline-block bg-green-100 text-green-700 text-xs font-bold tracking-widest uppercase px-3 py-1 rounded-full mb-1">
            Target & Objective
          </span>
          <h2 class="text-3xl font-black text-gray-900">K3 Target and Objective</h2>
        </div>
      </div>

      <p class="text-gray-400 text-sm mb-10 ml-16 anim-item">
        Key performance targets set to ensure occupational safety and health excellence across all operations.
      </p>

      <div class="space-y-3">
        <div v-for="(target, i) in k3Targets" :key="i"
          class="flex gap-4 bg-green-50 border border-green-100 rounded-xl p-5 hover:border-green-300 hover:bg-green-100/50 transition group anim-item"
          :style="{ animationDelay: (i * 0.08) + 's' }">
          <div class="w-9 h-9 rounded-xl bg-[#14532d] text-white font-black text-xs flex items-center justify-center flex-shrink-0 group-hover:bg-green-700 transition">
            {{ String(i + 1).padStart(2, '0') }}
          </div>
          <p class="text-gray-600 text-sm leading-relaxed self-center">{{ target }}</p>
        </div>
      </div>

    </div>
  </section>

  <!-- ── PROGRAM K3 ── -->
  <section class="py-20 bg-[#14532d]">
    <div class="max-w-5xl mx-auto px-10">

      <div class="flex items-center gap-4 mb-4 anim-item">
        <div class="w-12 h-12 rounded-2xl bg-green-500 flex items-center justify-center text-2xl flex-shrink-0">
          ⚙️
        </div>
        <div>
          <span class="inline-block bg-green-800 text-green-200 text-xs font-bold tracking-widest uppercase px-3 py-1 rounded-full mb-1">
            Program K3
          </span>
          <h2 class="text-3xl font-black text-white">K3 Programs</h2>
        </div>
      </div>

      <p class="text-green-200/60 text-sm leading-relaxed mb-10 ml-16 max-w-3xl anim-item">
        In achieving continuous improvement and ensuring that the system is well implemented, PT. Tanjungenim Lestari Pulp and Paper organizes K3 efforts that are contained in the K3 program. K3 program activities become a program that is integrated with the company's operational processes so that there is a mutually beneficial synergy without reducing one part. Some of the K3 programs carried out in 2019 are:
      </p>

      <div class="grid grid-cols-2 gap-4">
        <div v-for="(program, i) in k3Programs" :key="i"
          class="flex gap-4 bg-white/10 backdrop-blur-sm border border-white/10 rounded-xl p-5 hover:bg-white/15 transition group anim-item"
          :style="{ animationDelay: (i * 0.05) + 's' }">
          <div class="w-8 h-8 rounded-lg bg-green-500 text-white font-black text-xs flex items-center justify-center flex-shrink-0 group-hover:bg-green-400 transition">
            {{ String(i + 1).padStart(2, '0') }}
          </div>
          <p class="text-green-200/80 text-sm leading-relaxed self-center">{{ program }}</p>
        </div>
      </div>

      <!-- Slider Foto -->
      <div class="mt-12 relative rounded-2xl overflow-hidden shadow-xl aspect-video anim-item">
        <img
          v-for="(img, i) in programImages"
          :key="i"
          :src="img"
          alt="K3 Training Program"
          class="absolute inset-0 w-full h-full object-cover"
          :style="{ opacity: currentImg === i ? 1 : 0, transition: 'opacity 0.8s ease-in-out' }"
        />
        <div class="absolute inset-0 bg-gradient-to-t from-black/30 to-transparent"></div>

        <button @click="manualImgNav(prevImg)"
          class="absolute left-4 top-1/2 -translate-y-1/2 w-10 h-10 rounded-full bg-white/20 backdrop-blur-md text-white border border-white/30 flex items-center justify-center opacity-0 hover:opacity-100 transition-opacity cursor-pointer">
          ‹
        </button>
        <button @click="manualImgNav(nextImg)"
          class="absolute right-4 top-1/2 -translate-y-1/2 w-10 h-10 rounded-full bg-white/20 backdrop-blur-md text-white border border-white/30 flex items-center justify-center opacity-0 hover:opacity-100 transition-opacity cursor-pointer">
          ›
        </button>

        <div class="absolute bottom-3 left-1/2 -translate-x-1/2 flex gap-2 z-10">
          <button v-for="(_, i) in programImages" :key="i" @click="manualImgNav(() => currentImg = i)"
            class="h-1.5 rounded-full transition-all duration-500 cursor-pointer"
            :class="currentImg === i ? 'w-8 bg-white' : 'w-3 bg-white/40 hover:bg-white/60'">
          </button>
        </div>
      </div>
      <p class="text-center text-green-200/60 text-xs mt-3 italic anim-item">
        K3 Training & Emergency Response Program — PT. Tanjungenim Lestari Pulp and Paper
      </p>

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

</template>

<script setup>
import { ref, onMounted, onUnmounted, nextTick } from 'vue'
import { RouterLink } from 'vue-router'

const policyCards = [
  { icon: '📋', title: 'SMK3 & OHSAS 18001', desc: 'Certified management system ensuring continuous improvement and full compliance with government safety regulations.' },
  { icon: '🎓', title: 'Training & Resources', desc: 'All resources needed to ensure full compliance with Occupational Safety and Health requirements are provided.' },
  { icon: '🤝', title: 'All Parties Covered', desc: 'Policy covers all employees, contractors and suppliers operating within TEL factory areas.' },
]

const k3Targets = [
  'Update and adopt government regulations related to OSH by 100% in 2019.',
  'Renewed operating license deeds for all machines and OSH-related licenses for all equipment (lifts, lifting equipment, boilers and electricity) by 96% in 2019 issued by the Ministry of Manpower of the Republic of Indonesia.',
  'Ensure that companies, contractors, suppliers and vendors comply with regulations by 100% in 2019.',
  'Increase awareness of safety and safety in OHS risk assessments by providing training at least 10 times a year.',
  'Implement safety patrols a minimum of 10 times in one year.',
  'Conduct a hazard assessment in the work environment at least once a year.',
]

const k3Programs = [
  'Comply with government regulations related to occupational safety and health, both new and amended old regulations.',
  'Equip all work equipment and operational machinery with an operational permit.',
  'Ensuring companies, supplier contractors and vendors meet SMK3 requirements.',
  'Perform medical checkups.',
  'Conduct Training Need Analysis (TNA) related to K3 and conduct training for employees.',
  'K3 Promotion.',
  'Patrol and follow up findings / finding / K3 abnormality.',
  'Measure the Work Environment once a year and follow up abnormal parameters.',
  'Perform preventive maintenance for fire equipment.',
  'Implement work permits and fire work permits in the work environment.',
  'Investigating occupational accidents, Occupational Diseases (PAK) and fires.',
]

// ─── Image Slider ─────────────────────────────────────────────
const currentImg = ref(0)
let imgTimer = null

const programImages = [
  '/safety/2.png',
  '/safety/3.png',
  '/safety/4.png',
  '/safety/5.png',
  '/safety/6.png',
]

const nextImg = () => { currentImg.value = (currentImg.value + 1) % programImages.length }
const prevImg = () => { currentImg.value = (currentImg.value - 1 + programImages.length) % programImages.length }
const startImgTimer = () => { imgTimer = setInterval(nextImg, 5000) }
const stopImgTimer = () => { clearInterval(imgTimer) }
const manualImgNav = (fn) => { stopImgTimer(); fn(); startImgTimer() }

// ─── Scroll Observer ──────────────────────────────────────────
let observer = null

onMounted(async () => {
  await nextTick()
  initObserver()
  startImgTimer()
})

onUnmounted(() => {
  if (observer) {observer.disconnect()}
  stopImgTimer()
})

const initObserver = () => {
  observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.add('anim-visible')
      } else {
        entry.target.classList.remove('anim-visible')
      }
    })
  }, { 
    threshold: 0.15,
    rootMargin: '0px 0px -40px 0px'
  })

  document.querySelectorAll('.anim-item, .anim-slide-right, .anim-slide-left').forEach(el => {
    observer.observe(el)
  })
}
</script>

<style scoped>
@keyframes fadeUp {
  from { opacity: 0; transform: translateY(24px); }
  to { opacity: 1; transform: translateY(0); }
}
.anim-fade-up { animation: fadeUp 0.8s cubic-bezier(0.22, 1, 0.36, 1) forwards; }

@keyframes fadeRight {
  from { opacity: 0; transform: translateX(-24px); }
  to { opacity: 1; transform: translateX(0); }
}
.anim-fade-right { animation: fadeRight 0.8s cubic-bezier(0.22, 1, 0.36, 1) forwards; }

.anim-item {
  opacity: 0; transform: translateY(30px);
  transition: opacity 0.7s cubic-bezier(0.22, 1, 0.36, 1), transform 0.7s cubic-bezier(0.22, 1, 0.36, 1);
}
.anim-item.anim-visible { opacity: 1; transform: translateY(0); }

.anim-slide-right {
  opacity: 0; transform: translateX(-40px);
  transition: opacity 0.7s cubic-bezier(0.22, 1, 0.36, 1), transform 0.7s cubic-bezier(0.22, 1, 0.36, 1);
}
.anim-slide-right.anim-visible { opacity: 1; transform: translateX(0); }

.anim-slide-left {
  opacity: 0; transform: translateX(40px);
  transition: opacity 0.7s cubic-bezier(0.22, 1, 0.36, 1), transform 0.7s cubic-bezier(0.22, 1, 0.36, 1);
}
.anim-slide-left.anim-visible { opacity: 1; transform: translateX(0); }


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
</style>