# Rancher Turtles Release

This document describes the new release process for Turtles.

## Release Cadence

- New Rancher Turtles versions are released every month, as they follow Rancher's monthly release cadence. Since Turtles is bundled with Rancher, it needs to be tested and released before Rancher's tentative monthly release date. To facilitate this we typical release multiple release candidate (RC) versions of Turtles before Rancher's code freeze date, in order to give QA enough time to test its integration with Rancher and address potential issues early. Once QA has signed off an RC version of Turtles, we then release (un-RC) the signed off version and that is the version that will be released with Rancher in its upcoming new release.

## Release Process

We maintain 3 release branches for Turtles, one for each of the minor versions of Rancher that is under active maintenance, at any given point. This allows us to create new Turtles releases for each of the release branches as required for bug fixes and security patches. The process of cutting a new release is essentially the following:
- [Create a new tag for the release](#create-a-new-tag-for-the-release)
- [Update `rancher/charts` repository](#update-ranchercharts-repository)
- [Update `rancher/rancher` repository](#update-rancherrancher-repository)
- [Write release description for non-developers](#write-release-description-for-non-developers)
- [Rancher Turtles Community Documentation](#rancher-turtles-community-documentation) (required only when cutting minor/major versions)
- [Create a JIRA ticket for syncing Turtles Community Docs with Product Docs](#create-a-jira-ticket-for-syncing-turtles-community-docs-with-product-docs) (required only when cutting minor/major versions)

### Create a new tag for the release

Creating a new tag on `rancher/turtles` triggers a GitHub Actions [workflow](https://github.com/rancher/turtles/actions/workflows/release-v2.yaml) that builds the container images for the given tag and pushes them to their respective registries.

1. Clone the repository locally: 

```bash
git clone git@github.com:rancher/turtles.git
```

2. Depending on whether you are cutting a release candidate for a minor/major or patch version, the process varies:
  
  - If you are cutting a release candidate for a new minor/major:

      Create a new release branch (i.e. `release/v0.26`) and push it to the upstream repository:

      ```bash
        # Note: `upstream` must be the remote pointing to `github.com/rancher/turtles`.
        git checkout -b release/v0.26
        git push -u upstream release/v0.26
        # Export the tag of the minor/major release candidate to be cut, e.g.:
        export RELEASE_TAG='v0.26.0-rc.0'
      ```

  - If you are cutting a release candidate for a new patch:

      Use existing release branch:

      ```bash
        # Note: `upstream` must be the remote pointing to `github.com/rancher/turtles`.
        git checkout upstream/release/v0.26
        # Export the tag of the patch release candidate to be cut, e.g.:
        export RELEASE_TAG='v0.26.1-rc.0'
      ```

3. Create a signed/annotated tag for the new release from the release branch and push it:

  ```bash
  # Create tags locally
  git tag -s -a ${RELEASE_TAG} -m ${RELEASE_TAG}
  git tag -s -a test/${RELEASE_TAG} -m "Testing framework ${RELEASE_TAG}"
  git tag -s -a examples/${RELEASE_TAG} -m "ClusterClass examples ${RELEASE_TAG}"
  # Push tags
  # Note: `upstream` must be the remote pointing to `github.com/rancher/turtles`.
  git push upstream ${RELEASE_TAG}
  git push upstream test/${RELEASE_TAG}
  git push upstream examples/${RELEASE_TAG}
  ```

### Update `rancher/charts` repository

**Warning**: Before updating `rancher/charts` repository, ensure that the previous [step](#create-a-new-tag-for-the-release) has created a new GitHub release for the tag. This release must include the Helm chart archive file in its assets (typically named `rancher-turtles-${RELEASE_TAG}.tgz`).

This part of the release is automated via a GitHub Actions workflow that needs to be invoked manually from GitHub. To invoke it, navigate to the [workflow](https://github.com/rancher/turtles/actions/workflows/release-against-charts.yml), select the option `Run workflow` from the UI and pass the following parameters:
- Use workflow from: Branch: main
  This parameter should be set to the release branch that was used for creating the tag, for example `release/v0.25`. Using `main` branch may also work but using the release branch is safer, in case there are differences in the release workflows between these branches.
- Submit PR against the following rancher/charts branch (e.g. dev-v2.12): dev-v2.12
  This must be set to the `rancher/charts` branch that needs to be updated, with the new Turtles release. `dev-v2.12` is used for Rancher 2.12.x, `dev-v2.13` is used for Rancher 2.13.x and so on.
- Previous Turtles version (e.g. v0.23.0-rc.0)
  This is self explanatory, the value must be set to the previous Turtles version, for example `v0.25.1-rc.0`.
- New Turtles version (e.g. v0.23.0)
  This is self explanatory, the value must be set to the new Turtles version, for example `v0.25.1-rc.1`
- Set 'true' to bump chart major version when the Turtles minor version increases (e.g., v0.20.0 -> v0.21.0-rc.0). Default: false
  This is self explanatory, the values should be set to `true` when bumping the Turtles minor version, otherwise it should be set to `false`. 

Once this workflow has finished, a new PR should have been created in the `rancher/charts` repository that updates the selected branch with the new Turtles version. Here's an example (PR)[https://github.com/rancher/charts/pull/6294] from a previous run against the `dev-v2.13` branch. You need to review and merge this PR before proceeding to the next step.

### Update `rancher/rancher` repository

**Warning**: Before updating `rancher/rancher` repository, ensure that the PR generated from the previous [step](#update-ranchercharts-repository) has been merged.

This part of the release is also automated via a GitHub Actions workflow, that needs to be invoked manually from GitHub. To invoke it, navigate to the [workflow](https://github.com/rancher/turtles/actions/workflows/release-against-rancher.yml), select the option `Run workflow` from the UI and pass the following parameters:
- Use workflow from: Branch: main
  This parameter should be set to the release branch that was used for creating the tag, for example `release/v0.25`. Using `main` branch may also work but using the release branch is safer, in case there are differences in the release workflows between these branches.
- Submit PR against the following rancher/rancher branch (e.g. release/v2.12): release/v2.12
  This must be set to the `rancher/rancher` branch that needs to be updated, with the new Turtles release. `release/v2.12` is used for Rancher 2.12.x, `release/v2.13` (which is not yet present) will be used for Rancher 2.13.x and so on.
- Previous Turtles version (e.g. v0.23.0-rc.0)
  This is self explanatory, the value must be set to the previous Turtles version, for example `v0.25.1-rc.0`.
- New Turtles version (e.g. v0.23.0)
  This is self explanatory, the value must be set to the new Turtles version, for example `v0.25.1-rc.1`
- Set 'true' to bump chart major version when the Turtles minor version increases (e.g., v0.23.0 -> v0.24.0-rc.0). Default: false
  This is self explanatory, the values should be set to `true` when bumping the Turtles minor version, otherwise it should be set to `false`. 

Once this workflow has finished, a new PR should have been created in the `rancher/rancher` repository that updates the selected branch with the new Turtles version. You need to review and merge this PR. When this PR gets merged, you will have completed the process of releasing a new version of Turtles and including it in an upcoming version of Rancher.

### Write release description for non-developers

Turtles release notes are automatically generated based on the commits that are published with the new version.

These commit messages are usually hard to understand for those that are not very familiar with the project and we must add a "human-readable" description of what the release brings at the beginning of release notes. [This](https://github.com/rancher/turtles/releases/tag/v0.13.0) can be used for reference.

### Rancher Turtles Community Documentation

**Note**: This step is required only when cutting minor/major versions.

#### Update component versions in documentation

Before publishing a new version of the documentation, component versions used in the documentation must be updated.

1. Navigate to the Rancher Turtles Docs repository: [rancher/turtles-docs](https://github.com/rancher/turtles-docs)
2. Go to `Actions` and start a new run of `Update Component Versions`.
3. Look up the correct component versions from the [rancher/turtles](https://github.com/rancher/turtles) repository and fill them into the input fields. If any input fields are left blank, the workflow will retain the currently documented versions for those components.
4. This will generate a new PR in the same repository that updates the documentation with the specified versions.

#### Generate changelog from Turtles release notes

Before publishing a new version of the documentation, a new [Changelog](https://turtles.docs.rancher.com/turtles/next/en/changelogs/index.html) must be generated to include the new release.

1. Navigate to the Rancher Turtles Docs repository: [rancher/turtles-docs](https://github.com/rancher/turtles-docs)
2. Go to `Actions` and start a new run of `Updatecli`.
3. This will generate a new PR in the same repository that adds the release notes from any new Turtles release.

#### Publish Rancher Turtles Docs

1. Clone the Rancher Turtles Docs repository locally:

```bash
git clone git@github.com:rancher/turtles-docs.git
```

2. Export the tag of the minor/major release, create a signed/annotated tag and push it:

```bash
# Export the tag of the minor/major release in a format of v0.Y.Z/v1.Y.Z, e.g.:
export RELEASE_TAG=v0.26.0

# Create tags locally
git tag -s -a ${RELEASE_TAG} -m ${RELEASE_TAG}

# Push tags
# Note: `upstream` must be the remote pointing to `github.com/rancher/turtles-docs`
git push upstream ${RELEASE_TAG}
```

3. Wait for the [version publish workflow](https://github.com/rancher/turtles-docs/actions/workflows/version-publish.yaml) to create a pull request. The PR format is similar to [reference](https://github.com/rancher/turtles-docs/pull/160). Merging it would result in automatic documentation being published using the [publish workflow](https://github.com/rancher/turtles-docs/actions/workflows/publish.yaml).

The resulting state after the version publish workflow for the released tag is also stored under the `release-${RELEASE_TAG}` branch in the origin repository, which is consistent with the [branching strategy](#branches) for turtles.

Once all steps above are completed, a new version of Rancher Turtles Community Docs should be available at [https://turtles.docs.rancher.com].

### Create a JIRA ticket for syncing Turtles Community Docs with Product Docs

**Note**: This step is required only when cutting minor/major versions.

Follow the steps below to create a JIRA ticket:

- navigate to [JIRA](https://jira.suse.com/secure/Dashboard.jspa).
- take a look at reference [issue](https://jira.suse.com/browse/SURE-9171) and create a similar ticket for syncing the new Turtles Community Docs version that was published in [previous](#publish-a-new-rancher-turtles-community-docs-version) step with Turtles Product Docs.
In the ticket description, make sure to include the reference to latest published version of Rancher Turtles Community Docs and PR created automatically by GitHub Actions bot.

## Versioning

Rancher Turtles follows [semantic versioning](https://semver.org/) specification.

Example versions:
- Pre-release: `v0.4.0-rc.1`
- Minor release: `v0.4.0`
- Patch release: `v0.4.1`
- Major release: `v1.0.0`

With the v0 release of our codebase, we provide the following guarantees:

- A (*minor*) release CAN include:
  - Introduction of new API versions, or new Kinds.
  - Compatible API changes like field additions, deprecation notices, etc.
  - Breaking API changes for deprecated APIs, fields, or code.
  - Features, promotion or removal of feature gates.
  - And more!

- A (*patch*) release SHOULD only include backwards compatible set of bugfixes.

### Backporting

Any backport MUST not be breaking for either API or behavioral changes.

It is generally not accepted to submit pull requests directly against release branches (release/X). However, backports of fixes or changes that have already been merged into the main branch may be accepted to all supported branches:

- Critical bugs fixes, security issue fixes, or fixes for bugs without easy workarounds.
- Dependency bumps for CVE (usually limited to CVE resolution; backports of non-CVE related version bumps are considered exceptions to be evaluated case by case)
- Cert-manager version bumps (to avoid having releases with cert-manager versions that are out of support, when possible)
- Changes required to support new Kubernetes versions, when possible.
- Changes to use the latest Go patch version to build controller images.
- Improvements to existing docs (the latest supported branch hosts the current version of the book)

**Note:** We generally do not accept backports to Rancher Turtles release branches that are out of support.

#### Automation to create backports

There is automation in place to create backport PRs when a comment of a certain format is added to the original PR. For more details on how to do this, take a look at the [docs](https://github.com/rancher/turtles/blob/main/docs/release-automation-workflows.md#backport-pr-automation).

## Branches

Rancher Turtles has two types of branches: the `main` and `release/X` branches. Before integrating with Rancher release branches were named `release-X` but since then `release/X` is used.

The `main` branch is where development happens. All the latest and greatest code, including breaking changes, happens on main.

The `release/X` branches contain stable, backwards compatible code. On every major or minor release, a new branch is created. It is from these branches that minor, patch and pre-releases are tagged. In some cases, it may be necessary to open PRs for bugfixes directly against stable branches, but this should generally not be the case.

### Support and guarantees

Rancher Turtles maintains the most recent release/releases for all supported APIs. Support for this section refers to the ability to backport and release patch versions; [backport policy](#backporting) is defined above.

- The API version is determined from the GroupVersion defined in the top-level `api/` package.
- For the current stable API version (v1alpha1) we support the three most recent minor releases; older minor releases are immediately unsupported when a new major/minor release is available.
