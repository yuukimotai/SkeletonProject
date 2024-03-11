/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 */
const documents = {
    "query findAllMyBook {\n  findAllMyBook {\n    title\n    author\n    janCode\n    publisherName\n    itemUrl\n    largeImageUrl\n    mediumImageUrl\n    myBook\n  }\n}\n\nmutation createMyBookData($input: NewBook!) {\n  createMyBookData(input: $input) {\n    title\n    author\n    janCode\n  }\n}": types.FindAllMyBookDocument,
    "query findAllReadingMemo($input: GetAllMemo!) {\n  ReadingMemos(input: $input) {\n    janCode\n    title\n    content\n  }\n}\n\nmutation createReadingMemo($input: NewReadingMemo!) {\n  createReadingMemo(input: $input) {\n    janCode\n    title\n    content\n  }\n}\n\nmutation updateReadingMemo($input: UpdateReadingMemo!) {\n  updateReadingMemo(input: $input) {\n    memoId\n    janCode\n    title\n    content\n  }\n}\n\nmutation deleteReadingMemo($input: DeleteReadingMemo!) {\n  deleteReadingMemo(input: $input)\n}": types.FindAllReadingMemoDocument,
};

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = graphql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function graphql(source: string): unknown;

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "query findAllMyBook {\n  findAllMyBook {\n    title\n    author\n    janCode\n    publisherName\n    itemUrl\n    largeImageUrl\n    mediumImageUrl\n    myBook\n  }\n}\n\nmutation createMyBookData($input: NewBook!) {\n  createMyBookData(input: $input) {\n    title\n    author\n    janCode\n  }\n}"): (typeof documents)["query findAllMyBook {\n  findAllMyBook {\n    title\n    author\n    janCode\n    publisherName\n    itemUrl\n    largeImageUrl\n    mediumImageUrl\n    myBook\n  }\n}\n\nmutation createMyBookData($input: NewBook!) {\n  createMyBookData(input: $input) {\n    title\n    author\n    janCode\n  }\n}"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "query findAllReadingMemo($input: GetAllMemo!) {\n  ReadingMemos(input: $input) {\n    janCode\n    title\n    content\n  }\n}\n\nmutation createReadingMemo($input: NewReadingMemo!) {\n  createReadingMemo(input: $input) {\n    janCode\n    title\n    content\n  }\n}\n\nmutation updateReadingMemo($input: UpdateReadingMemo!) {\n  updateReadingMemo(input: $input) {\n    memoId\n    janCode\n    title\n    content\n  }\n}\n\nmutation deleteReadingMemo($input: DeleteReadingMemo!) {\n  deleteReadingMemo(input: $input)\n}"): (typeof documents)["query findAllReadingMemo($input: GetAllMemo!) {\n  ReadingMemos(input: $input) {\n    janCode\n    title\n    content\n  }\n}\n\nmutation createReadingMemo($input: NewReadingMemo!) {\n  createReadingMemo(input: $input) {\n    janCode\n    title\n    content\n  }\n}\n\nmutation updateReadingMemo($input: UpdateReadingMemo!) {\n  updateReadingMemo(input: $input) {\n    memoId\n    janCode\n    title\n    content\n  }\n}\n\nmutation deleteReadingMemo($input: DeleteReadingMemo!) {\n  deleteReadingMemo(input: $input)\n}"];

export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;