import {
  TechRadarApi,
  TechRadarLoaderResponse,
} from '@backstage-community/plugin-tech-radar';

export class ShortLinkTechRadar implements TechRadarApi {
  // @ts-ignore
  async load(id: string | undefined): Promise<TechRadarLoaderResponse> {
    // For example, this converts the timeline dates into date objects
    return {
      quadrants: [
        { id: 'languages', name: 'Languages' },
        { id: 'platforms', name: 'Platforms' },
        { id: 'tools', name: 'Tools' },
        { id: 'techniques', name: 'Techniques' },
      ],
      rings: [
        { id: "adopt", name: "Adopt", color: "#3884ff" },
        { id: "trial", name: "Trial", color: "#f9c80e" },
        { id: "assess", name: "Assess", color: "#f0932b" },
        { id: "hold", name: "Hold", color: "#eb4d4b" },
      ],
      // @ts-ignore
      entries: [
        {
          timeline: [
            {
              moved: 0,
              ringId: 'adopt',
              date: new Date('2020-08-06'),
              description:
                'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua',
            },
          ],
          url: '#',
          key: 'javascript',
          id: 'javascript',
          title: 'JavaScript',
          quadrant: 'languages',
          description:
            'Excepteur **sint** occaecat *cupidatat* non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.\n\n```ts\nconst x = "3";\n```\n',
        },
        {
          url: "#",
          key: "kubernetes",
          id: "kubernetes",
          title: "Kubernetes",
          quadrant: "platforms",
          description: "Kubernetes is an open-source system for automating deployment, scaling, and management of containerized applications.",
          timeline: [
            {
              moved: 0,
              ringId: "adopt",
              date: new Date("2020-08-06"),
              description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
            },
          ]
        },
        {
          title: "GitLab",
          quadrant: "tools",
          url: "#",
          key: "gitlab",
          id: "gitlab",
          description: "GitLab is a web-based DevOps lifecycle tool that provides a Git-repository manager providing wiki, issue-tracking and continuous integration and deployment pipeline features, using an open-source license, developed by GitLab Inc.",
          timeline: [
            {
              moved: 0,
              ringId: "adopt",
              date: new Date("2020-08-06"),
              description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
            },
          ],
        }
      ],
    };
  }
}
