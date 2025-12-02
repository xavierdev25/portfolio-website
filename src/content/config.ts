import { defineCollection, z } from 'astro:content';

const projectsCollection = defineCollection({
    type: 'content',
    schema: z.object({
        title: z.string(),
        category: z.string(),
        description: z.string(),
        images: z.array(z.string()),
        technologies: z.array(z.string()),
        repositoryUrl: z.string().optional(),
        liveUrl: z.string().optional(),
        order: z.number().default(0),
    }),
});

export const collections = {
    projects: projectsCollection,
};