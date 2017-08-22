package main

/* AUTOMATICALLY GENERATED FROM errors.yaml - DO NOT EDIT */

import (
	"fmt"
	"github.com/infobloxopen/themis/pdp-control"
	"strings"
)

const (
	externalErrorID = iota
	multiErrorID
	tracingTypeErrorID
	unknownEffectErrorID
	unknownAttributeTypeErrorID
	contextCreationErrorID
	policyCalculationErrorID
	effectTranslationErrorID
	effectCombiningErrorID
	obligationTranslationErrorID
	queueOverflowErrorID
	unknownUploadRequestErrorID
	invalidFromTagErrorID
	invalidToTagErrorID
	invalidTagsErrorID
	tagCheckErrorID
	emptyUploadErrorID
	unknownUploadErrorID
	policyUploadParseErrorID
	contentUploadParseErrorID
	missingPolicyStorageErrorID
	policyTransactionInProgressErrorID
	policyTransactionCreationErrorID
	policyUpdateParseErrorID
	policyUpdateApplicationErrorID
	policyTransactionCommitErrorID
	contentTransactionInProgressErrorID
	contentTransactionCreationErrorID
	contentUpdateParseErrorID
	contentUpdateApplicationErrorID
	contentTransactionCommitErrorID
)

type externalError struct {
	errorLink
	err error
}

func newExternalError(err error) *externalError {
	return &externalError{
		errorLink: errorLink{id: externalErrorID},
		err:       err}
}

func (e *externalError) Error() string {
	return e.errorf("%s", e.err)
}

type multiError struct {
	errorLink
	errs []error
}

func newMultiError(errs []error) *multiError {
	return &multiError{
		errorLink: errorLink{id: multiErrorID},
		errs:      errs}
}

func (e *multiError) Error() string {
	msgs := make([]string, len(e.errs))
	for i, err := range e.errs {
		msgs[i] = fmt.Sprintf("%q", err.Error())
	}
	msg := strings.Join(msgs, ", ")

	return e.errorf("multiple errors: %s", msg)
}

type tracingTypeError struct {
	errorLink
	t string
}

func newTracingTypeError(t string) *tracingTypeError {
	return &tracingTypeError{
		errorLink: errorLink{id: tracingTypeErrorID},
		t:         t}
}

func (e *tracingTypeError) Error() string {
	return e.errorf("Unknown tracing type %q", e.t)
}

type unknownEffectError struct {
	errorLink
	effect int
}

func newUnknownEffectError(effect int) *unknownEffectError {
	return &unknownEffectError{
		errorLink: errorLink{id: unknownEffectErrorID},
		effect:    effect}
}

func (e *unknownEffectError) Error() string {
	return e.errorf("Unknown policy effect %d", e.effect)
}

type unknownAttributeTypeError struct {
	errorLink
	t string
}

func newUnknownAttributeTypeError(t string) *unknownAttributeTypeError {
	return &unknownAttributeTypeError{
		errorLink: errorLink{id: unknownAttributeTypeErrorID},
		t:         t}
}

func (e *unknownAttributeTypeError) Error() string {
	return e.errorf("Unknown attribute type %q", e.t)
}

type contextCreationError struct {
	errorLink
	err error
}

func newContextCreationError(err error) *contextCreationError {
	return &contextCreationError{
		errorLink: errorLink{id: contextCreationErrorID},
		err:       err}
}

func (e *contextCreationError) Error() string {
	return e.errorf("Failed to create request context: %s", e.err)
}

type policyCalculationError struct {
	errorLink
	err error
}

func newPolicyCalculationError(err error) *policyCalculationError {
	return &policyCalculationError{
		errorLink: errorLink{id: policyCalculationErrorID},
		err:       err}
}

func (e *policyCalculationError) Error() string {
	return e.errorf("Failed to process request: %s", e.err)
}

type effectTranslationError struct {
	errorLink
	err error
}

func newEffectTranslationError(err error) *effectTranslationError {
	return &effectTranslationError{
		errorLink: errorLink{id: effectTranslationErrorID},
		err:       err}
}

func (e *effectTranslationError) Error() string {
	return e.errorf("Failed to translate effect: %s", e.err)
}

type effectCombiningError struct {
	errorLink
	err error
}

func newEffectCombiningError(err error) *effectCombiningError {
	return &effectCombiningError{
		errorLink: errorLink{id: effectCombiningErrorID},
		err:       err}
}

func (e *effectCombiningError) Error() string {
	return e.errorf("Failed to make failure effect: %s", e.err)
}

type obligationTranslationError struct {
	errorLink
	err error
}

func newObligationTranslationError(err error) *obligationTranslationError {
	return &obligationTranslationError{
		errorLink: errorLink{id: obligationTranslationErrorID},
		err:       err}
}

func (e *obligationTranslationError) Error() string {
	return e.errorf("Failed to translate obligations: %s", e.err)
}

type queueOverflowError struct {
	errorLink
	idx int32
}

func newQueueOverflowError(idx int32) *queueOverflowError {
	return &queueOverflowError{
		errorLink: errorLink{id: queueOverflowErrorID},
		idx:       idx}
}

func (e *queueOverflowError) Error() string {
	return e.errorf("Can't enqueue more than %d items", e.idx)
}

type unknownUploadRequestError struct {
	errorLink
	t control.Item_DataType
}

func newUnknownUploadRequestError(t control.Item_DataType) *unknownUploadRequestError {
	return &unknownUploadRequestError{
		errorLink: errorLink{id: unknownUploadRequestErrorID},
		t:         t}
}

func (e *unknownUploadRequestError) Error() string {
	return e.errorf("Unknown upload request type: %d", e.t)
}

type invalidFromTagError struct {
	errorLink
	tag string
	err error
}

func newInvalidFromTagError(tag string, err error) *invalidFromTagError {
	return &invalidFromTagError{
		errorLink: errorLink{id: invalidFromTagErrorID},
		tag:       tag,
		err:       err}
}

func (e *invalidFromTagError) Error() string {
	return e.errorf("Can't treat %q as current tag: %s", e.tag, e.err)
}

type invalidToTagError struct {
	errorLink
	tag string
	err error
}

func newInvalidToTagError(tag string, err error) *invalidToTagError {
	return &invalidToTagError{
		errorLink: errorLink{id: invalidToTagErrorID},
		tag:       tag,
		err:       err}
}

func (e *invalidToTagError) Error() string {
	return e.errorf("Can't treat %q as new tag: %s", e.tag, e.err)
}

type invalidTagsError struct {
	errorLink
	tag string
}

func newInvalidTagsError(tag string) *invalidTagsError {
	return &invalidTagsError{
		errorLink: errorLink{id: invalidTagsErrorID},
		tag:       tag}
}

func (e *invalidTagsError) Error() string {
	return e.errorf("Can't update from %q tag to no tag", e.tag)
}

type tagCheckError struct {
	errorLink
	err error
}

func newTagCheckError(err error) *tagCheckError {
	return &tagCheckError{
		errorLink: errorLink{id: tagCheckErrorID},
		err:       err}
}

func (e *tagCheckError) Error() string {
	return e.errorf("Failed tag check: %s", e.err)
}

type emptyUploadError struct {
	errorLink
}

func newEmptyUploadError() *emptyUploadError {
	return &emptyUploadError{
		errorLink: errorLink{id: emptyUploadErrorID}}
}

func (e *emptyUploadError) Error() string {
	return e.errorf("Empty upload")
}

type unknownUploadError struct {
	errorLink
	id int32
}

func newUnknownUploadError(id int32) *unknownUploadError {
	return &unknownUploadError{
		errorLink: errorLink{id: unknownUploadErrorID},
		id:        id}
}

func (e *unknownUploadError) Error() string {
	return e.errorf("Can't find upload request with id %d", e.id)
}

type policyUploadParseError struct {
	errorLink
	id  int32
	err error
}

func newPolicyUploadParseError(id int32, err error) *policyUploadParseError {
	return &policyUploadParseError{
		errorLink: errorLink{id: policyUploadParseErrorID},
		id:        id,
		err:       err}
}

func (e *policyUploadParseError) Error() string {
	return e.errorf("Failed to parse policy %d: %s", e.id, e.err)
}

type contentUploadParseError struct {
	errorLink
	id  int32
	err error
}

func newContentUploadParseError(id int32, err error) *contentUploadParseError {
	return &contentUploadParseError{
		errorLink: errorLink{id: contentUploadParseErrorID},
		id:        id,
		err:       err}
}

func (e *contentUploadParseError) Error() string {
	return e.errorf("Failed to parse content %d: %s", e.id, e.err)
}

type missingPolicyStorageError struct {
	errorLink
}

func newMissingPolicyStorageError() *missingPolicyStorageError {
	return &missingPolicyStorageError{
		errorLink: errorLink{id: missingPolicyStorageErrorID}}
}

func (e *missingPolicyStorageError) Error() string {
	return e.errorf("No any policy to update")
}

type policyTransactionInProgressError struct {
	errorLink
}

func newPolicyTransactionInProgressError() *policyTransactionInProgressError {
	return &policyTransactionInProgressError{
		errorLink: errorLink{id: policyTransactionInProgressErrorID}}
}

func (e *policyTransactionInProgressError) Error() string {
	return e.errorf("Can't start new policy update or completely reload policy while another update is in progress")
}

type policyTransactionCreationError struct {
	errorLink
	id   int32
	item *Item
	err  error
}

func newPolicyTransactionCreationError(id int32, item *Item, err error) *policyTransactionCreationError {
	return &policyTransactionCreationError{
		errorLink: errorLink{id: policyTransactionCreationErrorID},
		id:        id,
		item:      item,
		err:       err}
}

func (e *policyTransactionCreationError) Error() string {
	return e.errorf("Can't create transaction for policy update %d from tag %q to %q: %s", e.id, e.item.fromTag.String(), e.item.toTag.String(), e.err)
}

type policyUpdateParseError struct {
	errorLink
	id   int32
	item *Item
	err  error
}

func newPolicyUpdateParseError(id int32, item *Item, err error) *policyUpdateParseError {
	return &policyUpdateParseError{
		errorLink: errorLink{id: policyUpdateParseErrorID},
		id:        id,
		item:      item,
		err:       err}
}

func (e *policyUpdateParseError) Error() string {
	return e.errorf("Failed to parse update %d from tag %q to %q: %s", e.id, e.item.fromTag.String(), e.item.toTag.String(), e.err)
}

type policyUpdateApplicationError struct {
	errorLink
	id   int32
	item *Item
	err  error
}

func newPolicyUpdateApplicationError(id int32, item *Item, err error) *policyUpdateApplicationError {
	return &policyUpdateApplicationError{
		errorLink: errorLink{id: policyUpdateApplicationErrorID},
		id:        id,
		item:      item,
		err:       err}
}

func (e *policyUpdateApplicationError) Error() string {
	return e.errorf("Failed to apply update %d from tag %q to %q: %s", e.id, e.item.fromTag.String(), e.item.toTag.String(), e.err)
}

type policyTransactionCommitError struct {
	errorLink
	id   int32
	item *Item
	err  error
}

func newPolicyTransactionCommitError(id int32, item *Item, err error) *policyTransactionCommitError {
	return &policyTransactionCommitError{
		errorLink: errorLink{id: policyTransactionCommitErrorID},
		id:        id,
		item:      item,
		err:       err}
}

func (e *policyTransactionCommitError) Error() string {
	return e.errorf("Failed to commit transaction %d from tag %q to %q: %s", e.id, e.item.fromTag.String(), e.item.toTag.String(), e.err)
}

type contentTransactionInProgressError struct {
	errorLink
	id string
}

func newContentTransactionInProgressError(id string) *contentTransactionInProgressError {
	return &contentTransactionInProgressError{
		errorLink: errorLink{id: contentTransactionInProgressErrorID},
		id:        id}
}

func (e *contentTransactionInProgressError) Error() string {
	return e.errorf("Can't start new content update or completely reload content %q while another update of the content is in progress", e.id)
}

type contentTransactionCreationError struct {
	errorLink
	id   int32
	item *Item
	err  error
}

func newContentTransactionCreationError(id int32, item *Item, err error) *contentTransactionCreationError {
	return &contentTransactionCreationError{
		errorLink: errorLink{id: contentTransactionCreationErrorID},
		id:        id,
		item:      item,
		err:       err}
}

func (e *contentTransactionCreationError) Error() string {
	return e.errorf("Can't create transaction for content %q update %d from tag %q to %q: %s", e.item.id, e.id, e.item.fromTag.String(), e.item.toTag.String(), e.err)
}

type contentUpdateParseError struct {
	errorLink
	id   int32
	item *Item
	err  error
}

func newContentUpdateParseError(id int32, item *Item, err error) *contentUpdateParseError {
	return &contentUpdateParseError{
		errorLink: errorLink{id: contentUpdateParseErrorID},
		id:        id,
		item:      item,
		err:       err}
}

func (e *contentUpdateParseError) Error() string {
	return e.errorf("Failed to parse content %q update %d from tag %q to %q: %s", e.item.id, e.id, e.item.fromTag.String(), e.item.toTag.String(), e.err)
}

type contentUpdateApplicationError struct {
	errorLink
	id   int32
	item *Item
	err  error
}

func newContentUpdateApplicationError(id int32, item *Item, err error) *contentUpdateApplicationError {
	return &contentUpdateApplicationError{
		errorLink: errorLink{id: contentUpdateApplicationErrorID},
		id:        id,
		item:      item,
		err:       err}
}

func (e *contentUpdateApplicationError) Error() string {
	return e.errorf("Failed to apply content %q update %d from tag %q to %q: %s", e.item.id, e.id, e.item.fromTag.String(), e.item.toTag.String(), e.err)
}

type contentTransactionCommitError struct {
	errorLink
	id   int32
	item *Item
	err  error
}

func newContentTransactionCommitError(id int32, item *Item, err error) *contentTransactionCommitError {
	return &contentTransactionCommitError{
		errorLink: errorLink{id: contentTransactionCommitErrorID},
		id:        id,
		item:      item,
		err:       err}
}

func (e *contentTransactionCommitError) Error() string {
	return e.errorf("Failed to commit content %q transaction %d from tag %q to %q: %s", e.item.id, e.id, e.item.fromTag.String(), e.item.toTag.String(), e.err)
}
