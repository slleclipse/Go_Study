package analyzer

import "gochannel/main/chapter_6/errors"

func genError(errMsg string) error {
	return errors.NewCrawlerError(errors.ERROR_TYPE_ANANLYZER, errMsg)
}

func genParameterError(errMsg string) error {
	return errors.NewCrawlerErrorBy(errors.ERROR_TYPE_ANANLYZER, errors.NewIllegalParameterError(errMsg))
}
