import hashlib
import json
from pathlib import Path
import tempfile
import unittest
import zipfile

from publish import pack, Client


class PackageTests(unittest.TestCase):
    def test_repeatable_and_excludes_secrets(self):
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory) / 'docs'; root.mkdir()
            (root / 'index.md').write_text('# Hello', encoding='utf-8')
            (root / '.env').write_text('secret')
            (root / '.hidden').mkdir(); (root / '.hidden/key.md').write_text('secret')
            (root / 'docs.json').write_text(json.dumps({'schemaVersion': 1}))
            first, second = Path(directory) / 'a.zip', Path(directory) / 'b.zip'
            pack(root, first); pack(root, second)
            self.assertEqual(hashlib.sha256(first.read_bytes()).digest(), hashlib.sha256(second.read_bytes()).digest())
            with zipfile.ZipFile(first) as archive:
                self.assertEqual(archive.namelist(), ['docs.json', 'index.md'])

    def test_empty_directory_rejected_without_output(self):
        with tempfile.TemporaryDirectory() as directory:
            target = Path(directory) / 'docs.zip'
            with self.assertRaises(ValueError):
                pack(directory, target)
            self.assertFalse(target.exists())

    def test_credentials_not_sent_to_unsafe_origin(self):
        for origin in ['http://example.com', 'https://user:pass@example.com', 'https://example.com/path', 'https://example.com?token=x']:
            with self.subTest(origin=origin), self.assertRaises(ValueError):
                Client(origin)


if __name__ == '__main__':
    unittest.main()
