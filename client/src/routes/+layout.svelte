<script>
  import "../app.pcss";
  import "iconify-icon";
</script>

<nav class="flex items-center justify-between bg-zinc-300 p-4">
  <div class="flex items-center space-x-4">
    <a href="/" class="text-gray-800 hover:text-gray-600">Home</a>
    <a href="/members" class="text-gray-800 hover:text-gray-600">Members</a>
  </div>
  <div>
    <a href="/login" class="px-4 py-2 text-gray-800 hover:text-gray-600">
      <iconify-icon icon="material-symbols:login" />
    </a>
  </div>
</nav>

<slot />
