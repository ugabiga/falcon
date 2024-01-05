import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  JSON: { input: any; output: any; }
  Time: { input: any; output: any; }
};

export type Authentication = {
  __typename?: 'Authentication';
  createdAt: Scalars['Time']['output'];
  id: Scalars['ID']['output'];
  identifier: Scalars['String']['output'];
  provider: AuthenticationProvider;
  updatedAt: Scalars['Time']['output'];
  user: User;
  userID: Scalars['ID']['output'];
};

export enum AuthenticationProvider {
  Facebook = 'facebook',
  Google = 'google'
}

export type CreateTaskInput = {
  currency: Scalars['String']['input'];
  days: Scalars['String']['input'];
  hours: Scalars['String']['input'];
  params?: InputMaybe<Scalars['JSON']['input']>;
  size: Scalars['Float']['input'];
  symbol: Scalars['String']['input'];
  tradingAccountID: Scalars['ID']['input'];
  type: Scalars['String']['input'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createTask: Task;
  createTradingAccount: TradingAccount;
  updateTask: Task;
  updateTradingAccount: Scalars['Boolean']['output'];
  updateUser: User;
};


export type MutationCreateTaskArgs = {
  input: CreateTaskInput;
};


export type MutationCreateTradingAccountArgs = {
  exchange: Scalars['String']['input'];
  key: Scalars['String']['input'];
  name: Scalars['String']['input'];
  secret: Scalars['String']['input'];
};


export type MutationUpdateTaskArgs = {
  input: UpdateTaskInput;
  taskID: Scalars['ID']['input'];
  tradingAccountID: Scalars['ID']['input'];
};


export type MutationUpdateTradingAccountArgs = {
  exchange?: InputMaybe<Scalars['String']['input']>;
  id: Scalars['ID']['input'];
  key?: InputMaybe<Scalars['String']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
  secret?: InputMaybe<Scalars['String']['input']>;
};


export type MutationUpdateUserArgs = {
  input: UpdateUserInput;
};

export type Query = {
  __typename?: 'Query';
  taskHistoryIndex: TaskHistoryIndex;
  taskIndex?: Maybe<TaskIndex>;
  tradingAccountIndex: TradingAccountIndex;
  userIndex: UserIndex;
};


export type QueryTaskHistoryIndexArgs = {
  taskID: Scalars['ID']['input'];
  tradingAccountID: Scalars['ID']['input'];
};


export type QueryTaskIndexArgs = {
  tradingAccountID?: InputMaybe<Scalars['ID']['input']>;
};

export type Task = {
  __typename?: 'Task';
  createdAt: Scalars['Time']['output'];
  cron: Scalars['String']['output'];
  currency: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  isActive: Scalars['Boolean']['output'];
  nextExecutionTime?: Maybe<Scalars['Time']['output']>;
  params?: Maybe<Scalars['JSON']['output']>;
  size: Scalars['Float']['output'];
  symbol: Scalars['String']['output'];
  taskHistories?: Maybe<Array<TaskHistory>>;
  tradingAccount: TradingAccount;
  tradingAccountID: Scalars['ID']['output'];
  type: Scalars['String']['output'];
  updatedAt: Scalars['Time']['output'];
};

export type TaskHistory = {
  __typename?: 'TaskHistory';
  createdAt: Scalars['Time']['output'];
  id: Scalars['ID']['output'];
  isSuccess: Scalars['Boolean']['output'];
  task: Task;
  taskID: Scalars['ID']['output'];
  updatedAt: Scalars['Time']['output'];
};

export type TaskHistoryIndex = {
  __typename?: 'TaskHistoryIndex';
  task: Task;
  taskHistories?: Maybe<Array<TaskHistory>>;
};

export type TaskIndex = {
  __typename?: 'TaskIndex';
  selectedTradingAccount?: Maybe<TradingAccount>;
  tradingAccounts?: Maybe<Array<TradingAccount>>;
};

export type TradingAccount = {
  __typename?: 'TradingAccount';
  createdAt: Scalars['Time']['output'];
  exchange: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  ip: Scalars['String']['output'];
  key: Scalars['String']['output'];
  name: Scalars['String']['output'];
  tasks?: Maybe<Array<Task>>;
  updatedAt: Scalars['Time']['output'];
  user: User;
  userID: Scalars['ID']['output'];
};

export type TradingAccountIndex = {
  __typename?: 'TradingAccountIndex';
  tradingAccounts?: Maybe<Array<TradingAccount>>;
};

export type UpdateTaskInput = {
  currency: Scalars['String']['input'];
  days: Scalars['String']['input'];
  hours: Scalars['String']['input'];
  isActive: Scalars['Boolean']['input'];
  params?: InputMaybe<Scalars['JSON']['input']>;
  size: Scalars['Float']['input'];
  symbol: Scalars['String']['input'];
  type: Scalars['String']['input'];
};

export type UpdateUserInput = {
  name: Scalars['String']['input'];
  timezone: Scalars['String']['input'];
};

export type User = {
  __typename?: 'User';
  authentications?: Maybe<Array<Authentication>>;
  createdAt: Scalars['Time']['output'];
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  timezone: Scalars['String']['output'];
  tradingAccounts?: Maybe<Array<TradingAccount>>;
  updatedAt: Scalars['Time']['output'];
};

export type UserIndex = {
  __typename?: 'UserIndex';
  user: User;
};

/**
 * UserWhereInput is used for filtering User objects.
 * Input was generated by ent.
 */
export type UserWhereInput = {
  and?: InputMaybe<Array<UserWhereInput>>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** authentications edge predicates */
  hasAuthentications?: InputMaybe<Scalars['Boolean']['input']>;
  /** trading_accounts edge predicates */
  hasTradingAccounts?: InputMaybe<Scalars['Boolean']['input']>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  not?: InputMaybe<UserWhereInput>;
  or?: InputMaybe<Array<UserWhereInput>>;
  /** timezone field predicates */
  timezone?: InputMaybe<Scalars['String']['input']>;
  timezoneContains?: InputMaybe<Scalars['String']['input']>;
  timezoneContainsFold?: InputMaybe<Scalars['String']['input']>;
  timezoneEqualFold?: InputMaybe<Scalars['String']['input']>;
  timezoneGT?: InputMaybe<Scalars['String']['input']>;
  timezoneGTE?: InputMaybe<Scalars['String']['input']>;
  timezoneHasPrefix?: InputMaybe<Scalars['String']['input']>;
  timezoneHasSuffix?: InputMaybe<Scalars['String']['input']>;
  timezoneIn?: InputMaybe<Array<Scalars['String']['input']>>;
  timezoneIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  timezoneLT?: InputMaybe<Scalars['String']['input']>;
  timezoneLTE?: InputMaybe<Scalars['String']['input']>;
  timezoneNEQ?: InputMaybe<Scalars['String']['input']>;
  timezoneNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  timezoneNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
};

export type UserWithOptions = {
  withAuthentications: Scalars['Boolean']['input'];
  withTradingAccounts: Scalars['Boolean']['input'];
};

export type TaskFragmentFragment = { __typename?: 'Task', id: string, tradingAccountID: string, currency: string, size: number, symbol: string, cron: string, nextExecutionTime?: any | null, isActive: boolean, type: string, params?: any | null, updatedAt: any, createdAt: any };

export type GetTaskIndexQueryVariables = Exact<{
  tradingAccountID?: InputMaybe<Scalars['ID']['input']>;
}>;


export type GetTaskIndexQuery = { __typename?: 'Query', taskIndex?: { __typename?: 'TaskIndex', selectedTradingAccount?: { __typename?: 'TradingAccount', id: string, name: string, exchange: string, ip: string, key: string, tasks?: Array<{ __typename?: 'Task', id: string, tradingAccountID: string, currency: string, size: number, symbol: string, cron: string, nextExecutionTime?: any | null, isActive: boolean, type: string, params?: any | null, updatedAt: any, createdAt: any }> | null } | null, tradingAccounts?: Array<{ __typename?: 'TradingAccount', id: string, name: string, exchange: string, ip: string, key: string, tasks?: Array<{ __typename?: 'Task', id: string, tradingAccountID: string, currency: string, size: number, symbol: string, cron: string, nextExecutionTime?: any | null, isActive: boolean, type: string, params?: any | null, updatedAt: any, createdAt: any }> | null }> | null } | null };

export type CreateTaskMutationVariables = Exact<{
  tradingAccountID: Scalars['ID']['input'];
  currency: Scalars['String']['input'];
  size: Scalars['Float']['input'];
  symbol: Scalars['String']['input'];
  days: Scalars['String']['input'];
  hours: Scalars['String']['input'];
  type: Scalars['String']['input'];
  params?: InputMaybe<Scalars['JSON']['input']>;
}>;


export type CreateTaskMutation = { __typename?: 'Mutation', createTask: { __typename?: 'Task', id: string, tradingAccountID: string, currency: string, size: number, symbol: string, cron: string, nextExecutionTime?: any | null, isActive: boolean, type: string, params?: any | null, updatedAt: any, createdAt: any } };

export type UpdateTaskMutationVariables = Exact<{
  taskID: Scalars['ID']['input'];
  tradingAccountID: Scalars['ID']['input'];
  currency: Scalars['String']['input'];
  size: Scalars['Float']['input'];
  symbol: Scalars['String']['input'];
  days: Scalars['String']['input'];
  hours: Scalars['String']['input'];
  type: Scalars['String']['input'];
  params?: InputMaybe<Scalars['JSON']['input']>;
  isActive: Scalars['Boolean']['input'];
}>;


export type UpdateTaskMutation = { __typename?: 'Mutation', updateTask: { __typename?: 'Task', id: string, tradingAccountID: string, currency: string, size: number, symbol: string, cron: string, nextExecutionTime?: any | null, isActive: boolean, type: string, params?: any | null, updatedAt: any, createdAt: any } };

export type TaskHistoryFragmentFragment = { __typename?: 'TaskHistory', id: string, taskID: string, isSuccess: boolean, updatedAt: any, createdAt: any };

export type GetTaskHistoryIndexQueryVariables = Exact<{
  tradingAccountID: Scalars['ID']['input'];
  taskID: Scalars['ID']['input'];
}>;


export type GetTaskHistoryIndexQuery = { __typename?: 'Query', taskHistoryIndex: { __typename?: 'TaskHistoryIndex', task: { __typename?: 'Task', id: string, tradingAccountID: string, currency: string, size: number, symbol: string, cron: string, nextExecutionTime?: any | null, isActive: boolean, type: string, params?: any | null, updatedAt: any, createdAt: any }, taskHistories?: Array<{ __typename?: 'TaskHistory', id: string, taskID: string, isSuccess: boolean, updatedAt: any, createdAt: any }> | null } };

export type TradingAccountFragment = { __typename?: 'TradingAccount', id: string, userID: string, name: string, exchange: string, ip: string, key: string, updatedAt: any, createdAt: any };

export type TradingAccountIndexQueryVariables = Exact<{ [key: string]: never; }>;


export type TradingAccountIndexQuery = { __typename?: 'Query', tradingAccountIndex: { __typename?: 'TradingAccountIndex', tradingAccounts?: Array<{ __typename?: 'TradingAccount', id: string, userID: string, name: string, exchange: string, ip: string, key: string, updatedAt: any, createdAt: any }> | null } };

export type CreateTradingAccountMutationVariables = Exact<{
  name: Scalars['String']['input'];
  exchange: Scalars['String']['input'];
  key: Scalars['String']['input'];
  secret: Scalars['String']['input'];
}>;


export type CreateTradingAccountMutation = { __typename?: 'Mutation', createTradingAccount: { __typename?: 'TradingAccount', id: string, userID: string, name: string, exchange: string, ip: string, key: string, updatedAt: any, createdAt: any } };

export type UpdateTradingAccountMutationVariables = Exact<{
  id: Scalars['ID']['input'];
  name?: InputMaybe<Scalars['String']['input']>;
  exchange?: InputMaybe<Scalars['String']['input']>;
  key?: InputMaybe<Scalars['String']['input']>;
  secret?: InputMaybe<Scalars['String']['input']>;
}>;


export type UpdateTradingAccountMutation = { __typename?: 'Mutation', updateTradingAccount: boolean };

export type UserFragmentFragment = { __typename?: 'User', id: string, name: string, timezone: string };

export type UserIndexQueryVariables = Exact<{ [key: string]: never; }>;


export type UserIndexQuery = { __typename?: 'Query', userIndex: { __typename?: 'UserIndex', user: { __typename?: 'User', id: string, name: string, timezone: string } } };

export type UpdateUserMutationVariables = Exact<{
  name: Scalars['String']['input'];
  timezone: Scalars['String']['input'];
}>;


export type UpdateUserMutation = { __typename?: 'Mutation', updateUser: { __typename?: 'User', id: string, name: string, timezone: string } };

export const TaskFragmentFragmentDoc = {"kind":"Document","definitions":[{"kind":"FragmentDefinition","name":{"kind":"Name","value":"TaskFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Task"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"tradingAccountID"}},{"kind":"Field","name":{"kind":"Name","value":"currency"}},{"kind":"Field","name":{"kind":"Name","value":"size"}},{"kind":"Field","name":{"kind":"Name","value":"symbol"}},{"kind":"Field","name":{"kind":"Name","value":"cron"}},{"kind":"Field","name":{"kind":"Name","value":"nextExecutionTime"}},{"kind":"Field","name":{"kind":"Name","value":"isActive"}},{"kind":"Field","name":{"kind":"Name","value":"type"}},{"kind":"Field","name":{"kind":"Name","value":"params"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}}]}}]} as unknown as DocumentNode<TaskFragmentFragment, unknown>;
export const TaskHistoryFragmentFragmentDoc = {"kind":"Document","definitions":[{"kind":"FragmentDefinition","name":{"kind":"Name","value":"TaskHistoryFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"TaskHistory"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"taskID"}},{"kind":"Field","name":{"kind":"Name","value":"isSuccess"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}}]}}]} as unknown as DocumentNode<TaskHistoryFragmentFragment, unknown>;
export const TradingAccountFragmentDoc = {"kind":"Document","definitions":[{"kind":"FragmentDefinition","name":{"kind":"Name","value":"TradingAccount"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"TradingAccount"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"userID"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"exchange"}},{"kind":"Field","name":{"kind":"Name","value":"ip"}},{"kind":"Field","name":{"kind":"Name","value":"key"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}}]}}]} as unknown as DocumentNode<TradingAccountFragment, unknown>;
export const UserFragmentFragmentDoc = {"kind":"Document","definitions":[{"kind":"FragmentDefinition","name":{"kind":"Name","value":"UserFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"User"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"timezone"}}]}}]} as unknown as DocumentNode<UserFragmentFragment, unknown>;
export const GetTaskIndexDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"GetTaskIndex"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"tradingAccountID"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"taskIndex"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"tradingAccountID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"tradingAccountID"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"selectedTradingAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"exchange"}},{"kind":"Field","name":{"kind":"Name","value":"ip"}},{"kind":"Field","name":{"kind":"Name","value":"key"}},{"kind":"Field","name":{"kind":"Name","value":"tasks"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"TaskFragment"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"tradingAccounts"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"exchange"}},{"kind":"Field","name":{"kind":"Name","value":"ip"}},{"kind":"Field","name":{"kind":"Name","value":"key"}},{"kind":"Field","name":{"kind":"Name","value":"tasks"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"TaskFragment"}}]}}]}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"TaskFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Task"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"tradingAccountID"}},{"kind":"Field","name":{"kind":"Name","value":"currency"}},{"kind":"Field","name":{"kind":"Name","value":"size"}},{"kind":"Field","name":{"kind":"Name","value":"symbol"}},{"kind":"Field","name":{"kind":"Name","value":"cron"}},{"kind":"Field","name":{"kind":"Name","value":"nextExecutionTime"}},{"kind":"Field","name":{"kind":"Name","value":"isActive"}},{"kind":"Field","name":{"kind":"Name","value":"type"}},{"kind":"Field","name":{"kind":"Name","value":"params"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}}]}}]} as unknown as DocumentNode<GetTaskIndexQuery, GetTaskIndexQueryVariables>;
export const CreateTaskDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"CreateTask"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"tradingAccountID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"currency"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"size"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Float"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"symbol"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"days"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"hours"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"type"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"params"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"JSON"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createTask"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"tradingAccountID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"tradingAccountID"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"currency"},"value":{"kind":"Variable","name":{"kind":"Name","value":"currency"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"size"},"value":{"kind":"Variable","name":{"kind":"Name","value":"size"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"symbol"},"value":{"kind":"Variable","name":{"kind":"Name","value":"symbol"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"days"},"value":{"kind":"Variable","name":{"kind":"Name","value":"days"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"hours"},"value":{"kind":"Variable","name":{"kind":"Name","value":"hours"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"type"},"value":{"kind":"Variable","name":{"kind":"Name","value":"type"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"params"},"value":{"kind":"Variable","name":{"kind":"Name","value":"params"}}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"TaskFragment"}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"TaskFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Task"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"tradingAccountID"}},{"kind":"Field","name":{"kind":"Name","value":"currency"}},{"kind":"Field","name":{"kind":"Name","value":"size"}},{"kind":"Field","name":{"kind":"Name","value":"symbol"}},{"kind":"Field","name":{"kind":"Name","value":"cron"}},{"kind":"Field","name":{"kind":"Name","value":"nextExecutionTime"}},{"kind":"Field","name":{"kind":"Name","value":"isActive"}},{"kind":"Field","name":{"kind":"Name","value":"type"}},{"kind":"Field","name":{"kind":"Name","value":"params"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}}]}}]} as unknown as DocumentNode<CreateTaskMutation, CreateTaskMutationVariables>;
export const UpdateTaskDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"UpdateTask"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"taskID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"tradingAccountID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"currency"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"size"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Float"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"symbol"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"days"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"hours"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"type"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"params"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"JSON"}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"isActive"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Boolean"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"updateTask"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"taskID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"taskID"}}},{"kind":"Argument","name":{"kind":"Name","value":"tradingAccountID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"tradingAccountID"}}},{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"currency"},"value":{"kind":"Variable","name":{"kind":"Name","value":"currency"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"size"},"value":{"kind":"Variable","name":{"kind":"Name","value":"size"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"symbol"},"value":{"kind":"Variable","name":{"kind":"Name","value":"symbol"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"days"},"value":{"kind":"Variable","name":{"kind":"Name","value":"days"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"hours"},"value":{"kind":"Variable","name":{"kind":"Name","value":"hours"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"type"},"value":{"kind":"Variable","name":{"kind":"Name","value":"type"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"params"},"value":{"kind":"Variable","name":{"kind":"Name","value":"params"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"isActive"},"value":{"kind":"Variable","name":{"kind":"Name","value":"isActive"}}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"TaskFragment"}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"TaskFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Task"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"tradingAccountID"}},{"kind":"Field","name":{"kind":"Name","value":"currency"}},{"kind":"Field","name":{"kind":"Name","value":"size"}},{"kind":"Field","name":{"kind":"Name","value":"symbol"}},{"kind":"Field","name":{"kind":"Name","value":"cron"}},{"kind":"Field","name":{"kind":"Name","value":"nextExecutionTime"}},{"kind":"Field","name":{"kind":"Name","value":"isActive"}},{"kind":"Field","name":{"kind":"Name","value":"type"}},{"kind":"Field","name":{"kind":"Name","value":"params"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}}]}}]} as unknown as DocumentNode<UpdateTaskMutation, UpdateTaskMutationVariables>;
export const GetTaskHistoryIndexDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"GetTaskHistoryIndex"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"tradingAccountID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"taskID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"taskHistoryIndex"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"tradingAccountID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"tradingAccountID"}}},{"kind":"Argument","name":{"kind":"Name","value":"taskID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"taskID"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"task"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"TaskFragment"}}]}},{"kind":"Field","name":{"kind":"Name","value":"taskHistories"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"TaskHistoryFragment"}}]}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"TaskFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Task"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"tradingAccountID"}},{"kind":"Field","name":{"kind":"Name","value":"currency"}},{"kind":"Field","name":{"kind":"Name","value":"size"}},{"kind":"Field","name":{"kind":"Name","value":"symbol"}},{"kind":"Field","name":{"kind":"Name","value":"cron"}},{"kind":"Field","name":{"kind":"Name","value":"nextExecutionTime"}},{"kind":"Field","name":{"kind":"Name","value":"isActive"}},{"kind":"Field","name":{"kind":"Name","value":"type"}},{"kind":"Field","name":{"kind":"Name","value":"params"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"TaskHistoryFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"TaskHistory"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"taskID"}},{"kind":"Field","name":{"kind":"Name","value":"isSuccess"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}}]}}]} as unknown as DocumentNode<GetTaskHistoryIndexQuery, GetTaskHistoryIndexQueryVariables>;
export const TradingAccountIndexDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"TradingAccountIndex"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"tradingAccountIndex"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"tradingAccounts"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"TradingAccount"}}]}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"TradingAccount"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"TradingAccount"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"userID"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"exchange"}},{"kind":"Field","name":{"kind":"Name","value":"ip"}},{"kind":"Field","name":{"kind":"Name","value":"key"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}}]}}]} as unknown as DocumentNode<TradingAccountIndexQuery, TradingAccountIndexQueryVariables>;
export const CreateTradingAccountDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"CreateTradingAccount"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"name"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"exchange"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"key"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"secret"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createTradingAccount"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"name"},"value":{"kind":"Variable","name":{"kind":"Name","value":"name"}}},{"kind":"Argument","name":{"kind":"Name","value":"exchange"},"value":{"kind":"Variable","name":{"kind":"Name","value":"exchange"}}},{"kind":"Argument","name":{"kind":"Name","value":"key"},"value":{"kind":"Variable","name":{"kind":"Name","value":"key"}}},{"kind":"Argument","name":{"kind":"Name","value":"secret"},"value":{"kind":"Variable","name":{"kind":"Name","value":"secret"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"TradingAccount"}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"TradingAccount"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"TradingAccount"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"userID"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"exchange"}},{"kind":"Field","name":{"kind":"Name","value":"ip"}},{"kind":"Field","name":{"kind":"Name","value":"key"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}}]}}]} as unknown as DocumentNode<CreateTradingAccountMutation, CreateTradingAccountMutationVariables>;
export const UpdateTradingAccountDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"UpdateTradingAccount"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"name"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"exchange"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"key"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"secret"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"updateTradingAccount"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}},{"kind":"Argument","name":{"kind":"Name","value":"name"},"value":{"kind":"Variable","name":{"kind":"Name","value":"name"}}},{"kind":"Argument","name":{"kind":"Name","value":"exchange"},"value":{"kind":"Variable","name":{"kind":"Name","value":"exchange"}}},{"kind":"Argument","name":{"kind":"Name","value":"key"},"value":{"kind":"Variable","name":{"kind":"Name","value":"key"}}},{"kind":"Argument","name":{"kind":"Name","value":"secret"},"value":{"kind":"Variable","name":{"kind":"Name","value":"secret"}}}]}]}}]} as unknown as DocumentNode<UpdateTradingAccountMutation, UpdateTradingAccountMutationVariables>;
export const UserIndexDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"UserIndex"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"userIndex"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"UserFragment"}}]}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"UserFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"User"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"timezone"}}]}}]} as unknown as DocumentNode<UserIndexQuery, UserIndexQueryVariables>;
export const UpdateUserDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"UpdateUser"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"name"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"timezone"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"updateUser"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"name"},"value":{"kind":"Variable","name":{"kind":"Name","value":"name"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"timezone"},"value":{"kind":"Variable","name":{"kind":"Name","value":"timezone"}}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"UserFragment"}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"UserFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"User"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"timezone"}}]}}]} as unknown as DocumentNode<UpdateUserMutation, UpdateUserMutationVariables>;