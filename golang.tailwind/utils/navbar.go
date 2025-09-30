package utils

func GetNavbar() string {
	return `
        <nav class="fixed top-0 left-0 w-full bg-white border-b-1 border-gray-100 shadow-lg z-50">
        <div class="mx-auto max-w-7xl px-2 md:px-6 lg:px-8">
            <div class="relative flex h-16 items-center justify-between">

            <!-- Mobile menu button -->
            <div class="absolute inset-y-0 left-0 flex items-center md:hidden">
                <button id="menu-toggle"
                class="inline-flex items-center justify-center rounded-md p-2 text-gray-700 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-[#0265af]">
                <span class="sr-only">Open main menu</span>
                <!-- Icon: Hamburger -->
                <svg id="icon-open" class="block h-6 w-6" fill="none" stroke="currentColor" stroke-width="2"
                    viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round"
                    d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
                </svg>
                <!-- Icon: Close -->
                <svg id="icon-close" class="hidden h-6 w-6" fill="none" stroke="currentColor" stroke-width="2"
                    viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
                </svg>
                </button>
            </div>

            <!-- Logo + menu -->
            <div class="flex flex-1 items-center justify-center md:items-stretch md:justify-start">
                <div class="flex shrink-0 items-center">
                <img src="https://smkn4-web.netlify.app/_nuxt/Logo-SMK.Ba-Cc_BE.png" alt="Foto Sekolah"
            class="w-[35px] me-3">
                <h1 class="text-xl font-bold text-[#0265af]">SMKN 4 Tasikmalaya</h1>
                </div>
                <div class="hidden md:ml-6 md:block">
                <div class="flex space-x-4">
                    <a href="/beranda" class="px-3 py-2 rounded-md text-gray-700 hover:text-[#0265af]">Beranda</a>
                    <a href="/profil" class="px-3 py-2 rounded-md text-gray-700 hover:text-[#0265af]">Profil</a>
                    <a href="/ekstrakulikuler" class="px-3 py-2 rounded-md text-gray-700 hover:text-[#0265af]">Ekstrakulikuler</a>
                    <a href="/galeri" class="px-3 py-2 rounded-md text-gray-700 hover:text-[#0265af]">Galeri</a>
                </div>
                </div>
            </div>
            </div>
        </div>

        <!-- Mobile menu -->
        <div id="mobile-menu" class="hidden md:hidden transition-all duration-300 ease-in-out bg-white shadow-md">
            <div class="space-y-1 px-2 pb-3 pt-2">
            <a href="/beranda" class="block rounded-md px-3 py-2 text-gray-700 hover:bg-[#0265af] hover:text-white">Beranda</a>
            <a href="/profil" class="block rounded-md px-3 py-2 text-gray-700 hover:bg-[#0265af] hover:text-white">Profil</a>
            <a href="/ekstrakulikuler" class="block rounded-md px-3 py-2 text-gray-700 hover:bg-[#0265af] hover:text-white">Ekstrakulikuler</a>
            <a href="/galeri" class="block rounded-md px-3 py-2 text-gray-700 hover:bg-[#0265af] hover:text-white">Galeri</a>
            </div>
        </div>
        </nav>

        <script>
        const btn = document.getElementById('menu-toggle');
        const menu = document.getElementById('mobile-menu');
        const openIcon = document.getElementById('icon-open');
        const closeIcon = document.getElementById('icon-close');

        btn.addEventListener('click', () => {
            menu.classList.toggle('hidden');
            openIcon.classList.toggle('hidden');
            closeIcon.classList.toggle('hidden');
        });
        </script>

	`
}
