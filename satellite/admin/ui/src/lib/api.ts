// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

/*
 * This is the Admin API client exposing the API operations using the
 * interfaces, types and classes of the `ui-generator.ts` for allowing the
 * `UIGenerator.svelte` component to dynamically render their Web UI interface.
 */

import type { Operation } from '$lib/ui-generator';
import { InputText, Select } from '$lib/ui-generator';

// API must be implemented by any class which expose the access to a specific
// API.
export interface API {
	operations: {
		[key: string]: Operation[];
	};
}

export class Admin {
	readonly operations = {
		APIKeys: [
			{
				name: 'delete key',
				desc: 'Delete an API key',
				params: [['API key', new InputText('text', true)]],
				func: async (apiKey: string): Promise<null> => {
					return this.fetch('DELETE', `apikeys/${apiKey}`) as Promise<null>;
				}
			}
		],
		bucket: [
			{
				name: 'get',
				desc: 'Get the information of the specified bucket',
				params: [
					['Project ID', new InputText('text', true)],
					['Bucket name', new InputText('text', true)]
				],
				func: async (projectId: string, bucketName: string): Promise<Record<string, unknown>> => {
					return this.fetch('GET', `projects/${projectId}/buckets/${bucketName}`);
				}
			},
			{
				name: 'delete geofencing',
				desc: 'Delete the geofencing configuration of the specified bucket. The bucket MUST be empty',
				params: [
					['Project ID', new InputText('text', true)],
					['Bucket name', new InputText('text', true)]
				],
				func: async (projectId: string, bucketName: string): Promise<Record<string, unknown>> => {
					return this.fetch('DELETE', `projects/${projectId}/buckets/${bucketName}/geofence`);
				}
			},
			{
				name: 'set geofencing',
				desc: 'Set the geofencing configuration of the specified bucket. The bucket MUST be empty',
				params: [
					['Project ID', new InputText('text', true)],
					['Bucket name', new InputText('text', true)],
					[
						'Region',
						new Select(false, true, [
							{ text: 'European Union', value: 'EU' },
							{ text: 'European Economic Area', value: 'EEA' },
							{ text: 'United States', value: 'US' },
							{ text: 'Germany', value: 'DE' }
						])
					]
				],
				func: async (
					projectId: string,
					bucketName: string,
					region: string
				): Promise<Record<string, unknown>> => {
					const query = this.urlQueryFromObject({ region });
					if (query === '') {
						throw new APIError('region cannot be empty');
					}

					return this.fetch('POST', `projects/${projectId}/buckets/${bucketName}/geofence`, query);
				}
			}
		],
		project: [
			{
				name: 'create',
				desc: 'Add a new project to a specific user',
				params: [
					['Owner ID (user ID)', new InputText('text', true)],
					['Project Name', new InputText('text', true)]
				],
				func: async (ownerId: string, projectName: string): Promise<Record<string, unknown>> => {
					return this.fetch('POST', 'projects', null, { ownerId, projectName });
				}
			},
			{
				name: 'delete',
				desc: 'Delete a specific project',
				params: [['Project ID', new InputText('text', true)]],
				func: async (projectId: string): Promise<null> => {
					return this.fetch('DELETE', `projects/${projectId}`) as Promise<null>;
				}
			},
			{
				name: 'get',
				desc: 'Get the information of a specific project',
				params: [['Project ID', new InputText('text', true)]],
				func: async (projectId: string): Promise<Record<string, unknown>> => {
					return this.fetch('GET', `projects/${projectId}`);
				}
			},
			{
				name: 'update',
				desc: 'Update the information of a specific project',
				params: [
					['Project ID', new InputText('text', true)],
					['Project Name', new InputText('text', true)],
					['Description', new InputText('text', false)]
				],
				func: async (
					projectId: string,
					projectName: string,
					description: string
				): Promise<null> => {
					return this.fetch('PUT', `projects/${projectId}`, null, {
						projectName,
						description
					}) as Promise<null>;
				}
			},
			{
				name: 'create API key',
				desc: 'Create a new API key for a specific project',
				params: [
					['Project ID', new InputText('text', true)],
					['API key name', new InputText('text', true)]
				],
				func: async (projectId: string, name: string): Promise<Record<string, unknown>> => {
					return this.fetch('POST', `projects/${projectId}/apikeys`, null, {
						name
					});
				}
			},
			{
				name: 'delete API key',
				desc: 'Delete a API key of a specific project',
				params: [
					['Project ID', new InputText('text', true)],
					['API Key name', new InputText('text', true)]
				],
				func: async (projectId: string, apiKeyName: string): Promise<null> => {
					return this.fetch(
						'DELETE',
						`projects/${projectId}/apikeys/${apiKeyName}`
					) as Promise<null>;
				}
			},
			{
				name: 'get API keys',
				desc: 'Get the API keys of a specific project',
				params: [['Project ID', new InputText('text', true)]],
				func: async (projectId: string): Promise<Record<string, unknown>> => {
					return this.fetch('GET', `projects/${projectId}/apiKeys`);
				}
			},
			{
				name: 'get project usage',
				desc: 'Get the current usage of a specific project',
				params: [['Project ID', new InputText('text', true)]],
				func: async (projectId: string): Promise<Record<string, unknown>> => {
					return this.fetch('GET', `projects/${projectId}/usage`);
				}
			},
			{
				name: 'get project limits',
				desc: 'Get the current limits of a specific project',
				params: [['Project ID', new InputText('text', true)]],
				func: async (projectId: string): Promise<Record<string, unknown>> => {
					return this.fetch('GET', `projects/${projectId}/limit`);
				}
			},
			{
				name: 'update project limits',
				desc: 'Update the limits of a specific project',
				params: [
					['Project ID', new InputText('text', true)],
					['Storage (in bytes)', new InputText('number', false)],
					['Bandwidth (in bytes)', new InputText('number', false)],
					['Rate (requests per second)', new InputText('number', false)],
					['Buckets (maximum number)', new InputText('number', false)],
					['Segments (maximum number)', new InputText('number', false)]
				],
				func: async (
					projectId: string,
					usage: number,
					bandwidth: number,
					rate: number,
					buckets: number,
					segments: number
				): Promise<null> => {
					const query = this.urlQueryFromObject({
						usage,
						bandwidth,
						rate,
						buckets,
						segments
					});

					if (query === '') {
						throw new APIError('nothing to update, at least one limit must be set');
					}

					return this.fetch('PUT', `projects/${projectId}/limit`, query) as Promise<null>;
				}
			}
		],
		user: [
			{
				name: 'create',
				desc: 'Create a new user account',
				params: [
					['email', new InputText('email', true)],
					['full name', new InputText('text', false)],
					['password', new InputText('password', true)]
				],
				func: async (
					email: string,
					fullName: string,
					password: string
				): Promise<Record<string, unknown>> => {
					return this.fetch('POST', 'users', null, {
						email,
						fullName,
						password
					});
				}
			},
			{
				name: 'delete',
				desc: "Delete a user's account",
				params: [['email', new InputText('email', true)]],
				func: async (email: string): Promise<null> => {
					return this.fetch('DELETE', `users/${email}`) as Promise<null>;
				}
			},
			{
				name: 'get',
				desc: "Get the information of a user's account",
				params: [['email', new InputText('email', true)]],
				func: async (email: string): Promise<Record<string, unknown>> => {
					return this.fetch('GET', `users/${email}`);
				}
			},
			{
				name: 'update',
				desc: `Update the information of a user's account.
Blank fields will not be updated.`,
				params: [
					["current user's email", new InputText('email', true)],
					['new email', new InputText('email', false)],
					['full name', new InputText('text', false)],
					['short name', new InputText('text', false)],
					['partner ID', new InputText('text', false)],
					['password hash', new InputText('text', false)]
				],
				func: async (
					currentEmail: string,
					email?: string,
					fullName?: string,
					shortName?: string,
					partnerID?: string,
					passwordHash?: string
				): Promise<null> => {
					return this.fetch('PUT', `users/${currentEmail}`, null, {
						email,
						fullName,
						shortName,
						partnerID,
						passwordHash
					}) as Promise<null>;
				}
			}
		]
	};

	private readonly baseURL: string;

	constructor(baseURL: string, private readonly authToken: string) {
		this.baseURL = baseURL.endsWith('/') ? baseURL.substring(0, baseURL.length - 1) : baseURL;
	}

	protected async fetch(
		method: 'DELETE' | 'GET' | 'POST' | 'PUT',
		path: string,
		query?: string,
		data?: Record<string, unknown>
	): Promise<Record<string, unknown> | null> {
		const url = this.apiURL(path, query);
		const headers = new window.Headers({
			Authorization: this.authToken
		});

		let body: string;
		if (data) {
			headers.set('Content-Type', 'application/json');
			body = JSON.stringify(data);
		}

		const resp = await window.fetch(url, { method, headers, body });
		if (!resp.ok) {
			let body: Record<string, unknown>;
			if (resp.headers.get('Content-Type') === 'application/json') {
				body = await resp.json();
			}

			throw new APIError('server response error', resp.status, body);
		}

		if (resp.headers.get('Content-Type') === 'application/json') {
			return resp.json();
		}

		return null;
	}

	protected apiURL(path: string, query?: string): string {
		path = path.startsWith('/') ? path.substring(1) : path;

		if (!query) {
			query = '';
		} else {
			query = '?' + query;
		}

		return `${this.baseURL}/${path}${query}`;
	}

	protected urlQueryFromObject(values: Record<string, boolean | number | string>): string {
		const queryParts = [];

		for (const name of Object.keys(values)) {
			const val = values[name];
			if (val === undefined) {
				continue;
			}

			queryParts.push(`${name}=${encodeURIComponent(val)}`);
		}

		return queryParts.join('&');
	}
}

class APIError extends Error {
	constructor(
		public readonly msg: string,
		public readonly responseStatusCode?: number,
		public readonly responseBody?: Record<string, unknown> | string
	) {
		super(msg);
	}
}
