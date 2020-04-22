from distutils.core import setup
setup(
  name = 'pyodin',
  packages = ['pyodin'],
  version = '0.1.6',
  license='MIT',
  description = 'A python package for gathering information in odin jobs',
  url = 'https://gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7',
  author = 'James McDermott',
  author_email = 'james.mcdermott7@mail.dcu.ue',
  install_requires=[
          'ruamel.yaml',
          'pymongo',
      ],
  classifiers=[
    'Programming Language :: Python :: 3', 
    'Programming Language :: Python :: 3.5',
    'Programming Language :: Python :: 3.6',
    'Programming Language :: Python :: 3.7',
    'Programming Language :: Python :: 3.8',
  ],
)
