<script lang="ts">
  import { goto } from "$app/navigation";

  function oAuthGoogle() {
    // const redirectUri: string = 'http://localhost:8000/v1/auth/google';
    const redirectUri: string = "http://127.0.0.1:8000/v1/auth/google";
    goto(redirectUri);
  }
</script>

<div class="bg-tree_login bg-cover bg-no-repeat flex justify-center items-center h-screen">
  <div class="w-full max-w-xs">
    <form method="post" action="/login" class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
      <div class="mb-4">
        <input
          class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700
                     leading-tight focus:outline-none focus:shadow-outline"
          type="text"
          name="user"
          placeholder="Username"
        />
      </div>

      <div>
        <input
          class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3
                     leading-tight focus:outline-none focus:shadow-outline"
          type="password"
          name="password"
          placeholder="Password"
        />
      </div>

      <div class="flex justify-center">
        <button
          class="inline-flex h-10 w-full items-center justify-center gap-2 rounded border border-slate-300
                    bg-white p-2 text-sm font-medium text-black outline-none focus:ring-2 focus:ring-[#333] focus:ring-offset-1
                    disabled:cursor-not-allowed disabled:opacity-60"
          type="submit"
        >
          Sign in
        </button>
      </div>

      <div class="mt-6 text-center text-sm text-slate-600">
        <a href="/sign-up" class="font-medium text-zinc-500">Forgot password?</a>
      </div>

      <div class="flex w-full items-center gap-2 py-6 text-sm text-slate-600">
        <div class="h-px w-full bg-slate-200"></div>
        OR
        <div class="h-px w-full bg-slate-200"></div>
      </div>

      <div class="mt-3 flex flex-col gap-2">
        <button
          on:click={oAuthGoogle}
          class="inline-flex h-10 w-full items-center justify-center gap-2 rounded border border-slate-300 bg-white p-2 text-sm font-medium text-black outline-none focus:ring-2 focus:ring-[#333] focus:ring-offset-1 disabled:cursor-not-allowed disabled:opacity-60"
        >
          <img
            src="https://www.svgrepo.com/show/475656/google-color.svg"
            alt="Google"
            class="h-[18px] w-[18px]"
          />
          Continue with Google
        </button>
        <button
          class="inline-flex h-10 w-full items-center justify-center gap-2 rounded border border-slate-300 bg-white p-2 text-sm font-medium text-black outline-none focus:ring-2 focus:ring-[#333] focus:ring-offset-1 disabled:cursor-not-allowed disabled:opacity-60"
        >
          <img src="/instagram-1-svgrepo-com.svg" alt="Instagram" class="h-[18px] w-[18px]" />
          Continue with Instagram
        </button>
      </div>

      <div class="mt-6 text-center text-sm text-slate-600">
        Don't have an account?
        <a href="/sign-up" class="font-medium text-zinc-500">Sign up</a>
      </div>
    </form>
  </div>
</div>
