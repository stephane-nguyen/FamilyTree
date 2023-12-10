<script lang="ts">
  import { page } from "$app/stores";
  import { onMount } from "svelte";
  import type { Member } from "$lib/types";
  import { calculateAge } from "$lib/utils";

  let member: Member;
  let error: { message: any };
  const memberId: string = $page.params.memberId;
  const titles: string[] = ["Age", "Birthdate", "Gender", "Country", "City", "Description"];

  async function getMember() {
    try {
      const response = await fetch(`http://localhost:8000/v1/members/${memberId}`);

      if (response.ok) {
        return await response.json();
      } else {
        throw new Error("Failed to fetch member");
      }
    } catch (err) {
      throw err;
    }
  }

  onMount(() => {
    getMember()
      .then((data) => {
        member = data.member;
      })
      .catch((err) => {
        error = err;
      });
  });
</script>

{#if member}
  <div class="mt-5 ml-5">
    <div class="bg-white max-w-2xl shadow overflow-hidden sm:rounded-lg">
      <div class="px-4 py-5 sm:px-6">
        <h3 class="text-lg leading-6 font-medium text-gray-900">
          {member.firstname}
          {member.middlename}
          {member.lastname}
          <div class="float-right">
            <img src="/pen-square-svgrepo-com.svg" alt="edit" class="w-8 h-8 mr-3 inline" />
            <img src="/delete-profile-svgrepo-com.svg" alt="delete" class="w-8 h-8 inline" />
          </div>
        </h3>
        <p class="mt-1 max-w-2xl text-sm text-gray-500">Details and informations about user.</p>
      </div>
      <div class="border-t border-gray-200">
        <dl>
          <div class="bg-white px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
            {#each titles as title}
              <dt class="text-sm font-medium text-gray-500">
                {title}
              </dt>
              {#if title.toLowerCase() === "age"}
                <dd class="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
                  {calculateAge(member.birthdate)}
                </dd>
              {:else}
                <dd class="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
                  {#if member[title.toLowerCase()]}
                    {member[title.toLowerCase()]}
                  {/if}
                </dd>
              {/if}
            {/each}
          </div>
        </dl>
      </div>
    </div>
  </div>
{:else if error}
  <p style="color: red">{error.message}</p>
{:else}
  <p>Loading...</p>
{/if}
