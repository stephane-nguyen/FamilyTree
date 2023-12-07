<script lang="ts">
    import { page } from '$app/stores';
    import { onMount } from 'svelte';
    import type { Person } from '$lib/interfaces/Person';
    import { calculateAge } from '$lib/utils';


    let person: Person;
    let error: { message: any; };
    const memberId: string = $page.params.memberId;
    const titles: string[] = ["Age", "Birthdate", "Gender", "Country", "City", "Description"];
  
    async function getPerson() {
        try {
            const response = await fetch(`http://localhost:8000/v1/persons/${memberId}`);
    
            if (response.ok) {
                return await response.json();
            } else {
                throw new Error('Failed to fetch person');
            }
        } catch (err) {
            throw err;
        }
    }

    onMount(() => {
        getPerson()
            .then(data => {
                person = data.person;
            })
            .catch(err => {
                error = err;
            });
    });
  
</script>

{#if person}
    <div class="mt-5 ml-5">
        <div class="bg-white max-w-2xl shadow overflow-hidden sm:rounded-lg">
            <div class="px-4 py-5 sm:px-6">
                <h3 class="text-lg leading-6 font-medium text-gray-900">
                    {person.firstname} {person.middlename} {person.lastname}
                </h3>
                <p class="mt-1 max-w-2xl text-sm text-gray-500">
                    Details and informations about user.
                </p>
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
                                    {calculateAge(person.birthdate)}
                                </dd>
                            {:else}
                                <dd class="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
                                    {#if person[title.toLowerCase()]}
                                        {person[title.toLowerCase()]}
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
