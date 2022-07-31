package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func CopyFile(src string, dst string) error {
	content, err := ioutil.ReadFile(src)
	if err != nil {
		return fmt.Errorf("Failed to read %v: %w", src, err)
	}

	_, err = os.Stat(dst)
	if err == nil {
		// avoid overwriting if the destination file already contains the
		// desired content
		existingContent, err := ioutil.ReadFile(dst)
		if err == nil && bytes.Compare(existingContent, content) == 0 {
			return nil
		}

		err = os.Remove(dst)
		if err != nil {
			return fmt.Errorf("Failed to remove existing %v: %w", dst, err)
		}
	}

	statbuf, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("Failed to read %v: %w", src, err)
	}
	err = ioutil.WriteFile(dst, content, statbuf.Mode())
	if err != nil {
		return fmt.Errorf("Failed to write %v: %w", dst, err)
	}

	return nil
}

func CopyFileToDir(src string, dstDir string) error {
	dst := filepath.Join(dstDir, path.Base(src))
	return CopyFile(src, dst)
}

func CopyDir(src string, dst string) error {
	err := os.Mkdir(dst, 0755)
	if os.IsExist(err) {
		err = nil
	} else if err != nil {
		return err
	}

	entries, err := ioutil.ReadDir(src)
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		if entry.IsDir() {
			dstPath := filepath.Join(dst, entry.Name())
			err = CopyDir(srcPath, dstPath)
		} else {
			err = CopyFileToDir(srcPath, dst)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func SymlinkFile(src string, dst string) error {
	_, err := os.Stat(dst)
	if err == nil {
		// avoid overwriting if the destination file already contains the
		// desired content
		content, err := os.Readlink(dst)
		if err == nil && content == src {
			return nil
		}
		err = os.Remove(dst)
		if err != nil {
			return fmt.Errorf("Failed to remove existing %v: %w", dst, err)
		}
	}

	err = os.Symlink(src, dst)
	if err != nil {
		return fmt.Errorf("Failed to symlink %v: %w", dst, err)
	}

	return nil
}

func SymlinkFileToDir(src string, dstDir string) error {
	dst := filepath.Join(dstDir, path.Base(src))
	return SymlinkFile(src, dst)
}

func RenameFile(src string, dst string) error {
	removeSrcInsteadOfRename := false

	content, err := ioutil.ReadFile(src)
	if err != nil {
		return fmt.Errorf("Failed to read %v: %w", src, err)
	}

	_, err = os.Stat(dst)
	if err == nil {
		// avoid overwriting if the destination file already contains the
		// desired content
		existingContent, err := ioutil.ReadFile(dst)
		if err == nil && bytes.Compare(existingContent, content) == 0 {
			removeSrcInsteadOfRename = true
		}
	}

	if removeSrcInsteadOfRename {
		err = os.Remove(src)
		if err != nil {
			return fmt.Errorf("Failed to remove existing %v: %w", src, err)
		}
	} else {
		err = os.Rename(src, dst)
		if err != nil {
			return fmt.Errorf("Failed to rename %v -> %v: %w", src, dst, err)
		}
	}

	return nil
}

func RenameFileToDir(src string, dstDir string) error {
	dst := filepath.Join(dstDir, path.Base(src))
	return RenameFile(src, dst)
}
